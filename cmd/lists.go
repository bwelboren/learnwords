package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/gocolly/colly"
)

type Response struct {
	AuthToken string `json:"auth_token"`
}

type langList struct {
	ResultSections []struct {
		Name                       string `json:"name"`
		DisplayType                string `json:"display_type"`
		SearchType                 string `json:"search_type"`
		AppliedGradeFiltersDisplay string `json:"applied_grade_filters_display"`
		AppliedFilters             struct {
		} `json:"applied_filters"`
		AvailableFilters struct {
			SubjectID     []interface{} `json:"subject_id"`
			SchoolTrackID []struct {
				SectionName  string `json:"section_name"`
				SchoolTracks []struct {
					ID              int    `json:"id"`
					Name            string `json:"name"`
					EducationTypeID int    `json:"education_type_id"`
					Position        int    `json:"position"`
					Published       bool   `json:"published"`
				} `json:"school_tracks"`
			} `json:"school_track_id"`
			GradeNumber       []interface{} `json:"grade_number"`
			BookEditionNumber []interface{} `json:"book_edition_number"`
			BookChapter       []string      `json:"book_chapter"`
			BookParagraph     []string      `json:"book_paragraph"`
		} `json:"available_filters"`
		Results []struct {
			PublisherName       string `json:"publisher_name"`
			PublisherMethodName string `json:"publisher_method_name"`
			SchoolTracks        []struct {
				SchoolTrackID   int         `json:"school_track_id"`
				SchoolTrackName string      `json:"school_track_name"`
				GradeNumber     interface{} `json:"grade_number"`
			} `json:"school_tracks"`
			Subject struct {
				ID             interface{} `json:"id"`
				Name           string      `json:"name"`
				SpeechLocale   string      `json:"speech_locale"`
				SubjectLogoURL string      `json:"subject_logo_url"`
			} `json:"subject"`
			BookTitle        string      `json:"book_title"`
			BookEditionNr    interface{} `json:"book_edition_nr"`
			BookChapterNr    string      `json:"book_chapter_nr"`
			BookParagraphNr  string      `json:"book_paragraph_nr"`
			Isbn             interface{} `json:"isbn"`
			Title            string      `json:"title"`
			UpvoteCount      int         `json:"upvote_count"`
			DownvoteCount    int         `json:"downvote_count"`
			WordCount        int         `json:"word_count"`
			ListType         string      `json:"list_type"`
			GroupIds         []int       `json:"group_ids"`
			CreatorName      string      `json:"creator_name"`
			CreatorID        int         `json:"creator_id"`
			ID               string      `json:"id"`
			BookGradeDisplay string      `json:"book_grade_display"`
			IsFromPublisher  bool        `json:"is_from_publisher"`
			NetUpvotes       int         `json:"net_upvotes"`
			IsOwnList        bool        `json:"is_own_list"`
			ExerciseTypes    []struct {
				Code                string `json:"code"`
				Name                string `json:"name"`
				IconName            string `json:"icon_name"`
				Available           bool   `json:"available"`
				AvailabilityLabel   string `json:"availability_label,omitempty"`
				UnavailabilityLabel string `json:"unavailability_label,omitempty"`
			} `json:"exercise_types"`
			SubjectName    string `json:"subject_name"`
			SubjectLogoURL string `json:"subject_logo_url"`
			SubjectLocale  string `json:"subject_locale"`
		} `json:"results"`
		TotalCount int `json:"total_count"`
	} `json:"result_sections"`
}

type selectedList struct {
	ID                      int         `json:"id"`
	Title                   string      `json:"title"`
	Description             interface{} `json:"description"`
	Creator                 interface{} `json:"creator"`
	TimesPracticed          int         `json:"times_practiced"`
	Shared                  bool        `json:"shared"`
	Deleted                 bool        `json:"deleted"`
	PerformanceSectionOrder []string    `json:"performance_section_order"`
	WordsWithPerformance    []struct {
		Words                 []string `json:"words"`
		AllResultsPerformance struct {
			CorrectAnswers        int    `json:"correct_answers"`
			IncorrectAnswers      int    `json:"incorrect_answers"`
			CorrectnessPercentage int    `json:"correctness_percentage"`
			PerformanceCategory   string `json:"performance_category"`
		} `json:"all_results_performance"`
		Latest3ResultsPerformance struct {
			CorrectAnswers        int    `json:"correct_answers"`
			IncorrectAnswers      int    `json:"incorrect_answers"`
			CorrectnessPercentage int    `json:"correctness_percentage"`
			PerformanceCategory   string `json:"performance_category"`
		} `json:"latest_3_results_performance"`
		Subjects []struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Locale string `json:"locale"`
			Flag   string `json:"flag"`
		} `json:"subjects"`
		Locales []struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			LanguageCode string `json:"language_code"`
			FlagCode     string `json:"flag_code"`
			Code         string `json:"code"`
			Locale       string `json:"locale"`
			PractiseType string `json:"practise_type"`
		} `json:"locales"`
	} `json:"words_with_performance"`
	Book struct {
		ID               int    `json:"id"`
		Name             string `json:"name"`
		BookGradeDisplay string `json:"book_grade_display"`
		CoverPictureURL  string `json:"cover_picture_url"`
		SubjectIconURL   string `json:"subject_icon_url"`
		SubjectLogoURL   string `json:"subject_logo_url"`
	} `json:"book"`
	ExerciseTypes []struct {
		Code                string `json:"code"`
		Name                string `json:"name"`
		IconName            string `json:"icon_name"`
		Available           bool   `json:"available"`
		AvailabilityLabel   string `json:"availability_label,omitempty"`
		UnavailabilityLabel string `json:"unavailability_label,omitempty"`
	} `json:"exercise_types"`
	LiveBattleExerciseTypes []struct {
		Code                string `json:"code"`
		Name                string `json:"name"`
		IconName            string `json:"icon_name"`
		Available           bool   `json:"available"`
		UnavailabilityLabel string `json:"unavailability_label,omitempty"`
	} `json:"live_battle_exercise_types"`
	RelatedTopics  []interface{} `json:"related_topics"`
	PausedExercise struct {
		ID                 int `json:"id"`
		ProgressPercentage int `json:"progress_percentage"`
		ExerciseType       struct {
			Code              string `json:"code"`
			Name              string `json:"name"`
			IconName          string `json:"icon_name"`
			Available         bool   `json:"available"`
			AvailabilityLabel string `json:"availability_label"`
		} `json:"exercise_type"`
		WordQueue struct {
		} `json:"word_queue"`
	} `json:"paused_exercise"`
	Status            string      `json:"status"`
	NeedsUpgrade      bool        `json:"needs_upgrade"`
	MinRequiredRole   interface{} `json:"min_required_role"`
	WordCount         int         `json:"word_count"`
	UpgradeInfo       interface{} `json:"upgrade_info"`
	RelatedTopicsType string      `json:"related_topics_type"`
	Chapter           struct {
		ID           int         `json:"id"`
		Title        interface{} `json:"title"`
		TitleNumber  string      `json:"title_number"`
		TitleDisplay string      `json:"title_display"`
	} `json:"chapter"`
	Subjects []struct {
		ID       int    `json:"id"`
		Code     string `json:"code"`
		Flag     string `json:"flag"`
		Language struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Code     string `json:"code"`
			FlagCode string `json:"flag_code"`
		} `json:"language"`
		Name          string `json:"name"`
		Locale        string `json:"locale"`
		LocalizedName string `json:"localized_name"`
		Default       bool   `json:"default"`
	} `json:"subjects"`
	Subject struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Locale struct {
			ID       int    `json:"id"`
			Code     string `json:"code"`
			Flag     string `json:"flag"`
			Language struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Code     string `json:"code"`
				FlagCode string `json:"flag_code"`
			} `json:"language"`
			Name          string `json:"name"`
			Locale        string `json:"locale"`
			LocalizedName string `json:"localized_name"`
			Default       bool   `json:"default"`
		} `json:"locale"`
		IconURL  string `json:"icon_url"`
		Fallback bool   `json:"fallback"`
	} `json:"subject"`
	IsOwnList           bool `json:"is_own_list"`
	IsFromPublisher     bool `json:"is_from_publisher"`
	PrioritizedLanguage struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Locale string `json:"locale"`
		Flag   string `json:"flag"`
	} `json:"prioritized_language"`
	Locales []struct {
		ID       int    `json:"id"`
		Code     string `json:"code"`
		Flag     string `json:"flag"`
		Language struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Code     string `json:"code"`
			FlagCode string `json:"flag_code"`
		} `json:"language"`
		Name          string `json:"name"`
		Locale        string `json:"locale"`
		LocalizedName string `json:"localized_name"`
		Default       bool   `json:"default"`
	} `json:"locales"`
}

type WordList struct {
	ID    string
	Name  string
	Lang  []string
	Words []string
}

var (
	authToken string
	IDs       []string
)

func SetAuthToken(email, password string) {

	var resp Response

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("content-type", "application/json; charset=utf-8")
	})
	c.OnResponse(func(r *colly.Response) {
		if err := json.Unmarshal(r.Body, &resp); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
	})

	err := c.PostRaw("https://api.wrts.nl/api/v3/auth/get_token", []byte(fmt.Sprintf(`{"email":"%s","password":"%s"}`, email, password)))

	if err != nil {
		log.Fatal(err)
	}

	authToken = resp.AuthToken
}

func GetOfficialLists(language string) {

	var officialList langList

	url := fmt.Sprintf("https://api.wrts.nl/api/v3/search?apply_default_filters=true&search_terms=%s&limit=100&offset=0&type=official_lists", language)

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("x-auth-token", authToken)
	})

	c.OnResponse(func(r *colly.Response) {
		if err := json.Unmarshal(r.Body, &officialList); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}

		officialList.sortListIDs()
	})

	c.Visit(url)

}

func (l langList) sortListIDs() {

	lists := l.ResultSections[0].Results

	// sort by id
	sort.Slice(lists, func(i, j int) bool {
		return lists[i].ID < lists[j].ID
	})

	for list := range lists {
		IDs = append(IDs, lists[list].ID)
	}

}

func GetWordsByListID(id string) WordList {

	var sl selectedList

	myWords := []string{}

	url := fmt.Sprintf("https://api.wrts.nl/api/v3/public/lists/%s", id)

	lang := []string{}

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("x-auth-token", authToken)
	})

	c.OnResponse(func(r *colly.Response) {
		if err := json.Unmarshal(r.Body, &sl); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}

		// Working with one list

		leftColumnLanguage := sl.Subjects[0].Language.Name
		rightColumnLanguage := sl.Subjects[1].Language.Name
		lang = append(lang, leftColumnLanguage)
		lang = append(lang, rightColumnLanguage)

		for _, wordPair := range sl.WordsWithPerformance {
			myWords = append(myWords, wordPair.Words...)
		}

	})

	c.Visit(url)

	return WordList{ID: id, Name: sl.Title, Lang: lang, Words: myWords}

}

func GetAllWordsFromWordLists() []WordList {
	var wl []WordList

	for _, x := range IDs {
		wl = append(wl, GetWordsByListID(x))
	}

	return wl
}
