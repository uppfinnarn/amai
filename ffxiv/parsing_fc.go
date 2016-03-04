package ffxiv

import (
	"fmt"
	"strconv"
	"strings"
	"regexp"
	"github.com/PuerkitoBio/goquery"
)

func parseBanner(s string) (gc, server string, err error) {
	re := regexp.MustCompile(`\s*([\w\s]+)[^\(]+\(([^\)]+)\)`)
	matches := re.FindStringSubmatch(s)
	
	if len(matches) == 0 {
		return "", "", ConfusedByMarkupError("FC Banner regex didn't match")
	}
	
	return strings.TrimSpace(matches[1]), strings.TrimSpace(matches[2]), nil
}

func parseFreeCompany(id string, doc *goquery.Document) (fc FFXIVFreeCompany, err error) {
	fc = FFXIVFreeCompany{}
	
	fc.ID, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		return fc, err
	}
	
	banner := doc.Find(".ic_freecompany_box").Text()
	fc.GrandCompany, fc.Server, err = parseBanner(banner)
	if err != nil {
		return fc, err
	}
	
	nameTagRE := regexp.MustCompile(`([^«]+)\s*«([^»]+)»`)
	doc.Find(".area_inner_body tr").EachWithBreak(func(i int, e *goquery.Selection) bool {
		key := e.Find("th").Text()
		valE := e.Find("td")
		
		switch key {
		case "Free Company Name«Company Tag»":
			nameTagMatches := nameTagRE.FindStringSubmatch(valE.Text())
			if len(nameTagMatches) == 0 {
				err = ConfusedByMarkupError("Can't parse FC name/tag")
				return false
			}
			fc.Name = strings.TrimSpace(nameTagMatches[1])
			fc.Tag = strings.TrimSpace(nameTagMatches[2])
		case "Formed":
		case "Active Members":
			// Skipping this in favor of parsing the full member list.
		case "Rank":
			fc.Rank, err = strconv.Atoi(strings.TrimSpace(valE.Text()))
		case "Ranking":
			// Rather uninteresting, purely ephemeral information; could parse
			// this if The Feast makes it interesting, I suppose? I honestly
			// don't even understand what's graded here.
		case "Company Slogan":
			fc.Description, err = valE.Html()
			fc.Description = strings.Replace(fc.Description, "<br/>", "\n", -1)
		case "Focus":
			lis := e.Find("li")
			if lis.Length() == 0 {
				err = ConfusedByMarkupError("There are no focus li's!")
				return false
			}
			
			lis.EachWithBreak(func(j int, li *goquery.Selection) bool {
				state := !li.HasClass("icon_off")
				focus := li.Find("img").AttrOr("title", "")
				switch focus {
				case "Role-playing": fc.Focus.RolePlaying = state
				case "Leveling": fc.Focus.Leveling = state
				case "Casual": fc.Focus.Casual = state
				case "Hardcore": fc.Focus.Hardcore = state
				case "Dungeons": fc.Focus.Dungeons = state
				case "Guildhests": fc.Focus.Guildhests = state
				case "Trials": fc.Focus.Trials = state
				case "Raids": fc.Focus.Raids = state
				case "PvP": fc.Focus.PvP = state
				default:
					err = ConfusedByMarkupError(fmt.Sprintf("Unknown focus: %s", focus))
				}
				return err == nil
			})
		case "Seeking":
			lis := e.Find("li")
			if lis.Length() == 0 {
				err = ConfusedByMarkupError("There are no seeking li's!")
				return false
			}
			
			lis.EachWithBreak(func(j int, li *goquery.Selection) bool {
				state := !li.HasClass("icon_off")
				focus := li.Find("img").AttrOr("title", "")
				switch focus {
				case "Tank": fc.Seeking.Tank = state
				case "Healer": fc.Seeking.Healer = state
				case "DPS": fc.Seeking.DPS = state
				case "Crafter": fc.Seeking.Crafter = state
				case "Gatherer": fc.Seeking.Gatherer = state
				default:
					err = ConfusedByMarkupError(fmt.Sprintf("Unknown seeking: %s", focus))
				}
				return err == nil
			})
		case "Active":
			// Not sure what the values here are... I'll have a look ingame later.
		case "Recruitment":
		case "Estate Profile":
		default:
			err = ConfusedByMarkupError(fmt.Sprintf("Unknown item: %s", key))
		}
		return err == nil
	})
	if err != nil {
		return fc, err
	}
	
	return fc, nil
}
