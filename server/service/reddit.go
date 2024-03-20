package service

import (
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "net/http"
  "os"
  "strings"
)
var token string
type RedditService struct {
  Url string
  Payload *strings.Reader
  Method string
  Client *http.Client
}

func(r RedditService) FetchDataFromReddit(subreddit string) []byte {
  r.Method = "GET"
  r.Url = fmt.Sprintf("https://oauth.reddit.com/%s/about", subreddit)
  req, err := http.NewRequest(r.Method, r.Url, nil)
  if err != nil {
    fmt.Println(err)
    return nil
  }
  req.Header.Add("Authorization", token)
  req.Header.Add("User-Agent", "PostmanRuntime/7.36.1")

  res, err := r.Client.Do(req)

  if err != nil {
    fmt.Println(err)
    return nil
  }
  defer res.Body.Close()

  if b, err := io.ReadAll(res.Body); err == nil {
    return b
  }
  fmt.Println(err)
  return nil
}

func(r RedditService) RefreshToken() {
  r.Method = "POST"
  r.Url = "https://www.reddit.com/api/v1/access_token"
  r.Payload = strings.NewReader(fmt.Sprintf("username=%s&password=%s&grant_type=password", os.Getenv("REDDIT_USER"), os.Getenv("REDDIT_PWD")))

  req, err := http.NewRequest(r.Method, r.Url, r.Payload)
  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Authorization", os.Getenv("REDDIT_BASIC_AUTH"))

  res, err := r.Client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  var responseData map[string]interface{}
  err = json.NewDecoder(res.Body).Decode(&responseData)
  accessToken, ok := responseData["access_token"]
  if ok {
    token = "bearer " + accessToken.(string)
  }
}

func(r RedditService) UnMarshalAboutResponseToStruct(res []byte, rs *SubRedditAboutResponse) error {
  var er = &SubRedditAboutErrorResponse{}
  json.Unmarshal(res, er)
  if er.Error == 429 {
    fmt.Println("error: ", string(res[:]), er.Error)
    return SubRedditAboutErrorResponseTooMany
  }else if er.Error >= 400 {
    fmt.Println("error: ", string(res[:]), er.Error)
    return SubRedditAboutErrorResponseBanned
  }
  json.Unmarshal(res, &rs)
  return nil
}

type SubRedditAboutErrorResponse struct {
  Message string `json:"message"`
  Error float64 `json:"error"`
  Reason string `json:"reason"`
}

var SubRedditAboutErrorResponseBanned = errors.New("Subreddit banned")
var SubRedditAboutErrorResponseTooMany = errors.New("Too many request")

type SubRedditAboutResponse struct {
  Kind string `json:"kind"`
  Data struct {
    UserFlairBackgroundColor       any    `json:"user_flair_background_color"`
    SubmitTextHTML                 string `json:"submit_text_html"`
    RestrictPosting                bool   `json:"restrict_posting"`
    UserIsBanned                   bool   `json:"user_is_banned"`
    FreeFormReports                bool   `json:"free_form_reports"`
    WikiEnabled                    bool   `json:"wiki_enabled"`
    UserIsMuted                    bool   `json:"user_is_muted"`
    UserCanFlairInSr               bool   `json:"user_can_flair_in_sr"`
    DisplayName                    string `json:"display_name"`
    HeaderImg                      any    `json:"header_img"`
    Title                          string `json:"title"`
    AllowGalleries                 bool   `json:"allow_galleries"`
    IconSize                       []int  `json:"icon_size"`
    PrimaryColor                   string `json:"primary_color"`
    ActiveUserCount                int    `json:"active_user_count"`
    IconImg                        string `json:"icon_img"`
    DisplayNamePrefixed            string `json:"display_name_prefixed"`
    AccountsActive                 int    `json:"accounts_active"`
    PublicTraffic                  bool   `json:"public_traffic"`
    Subscribers                    int    `json:"subscribers"`
    UserFlairRichtext              []any  `json:"user_flair_richtext"`
    VideostreamLinksCount          int    `json:"videostream_links_count"`
    Name                           string `json:"name"`
    Quarantine                     bool   `json:"quarantine"`
    HideAds                        bool   `json:"hide_ads"`
    PredictionLeaderboardEntryType int    `json:"prediction_leaderboard_entry_type"`
    EmojisEnabled                  bool   `json:"emojis_enabled"`
    AdvertiserCategory             string `json:"advertiser_category"`
    PublicDescription              string `json:"public_description"`
    CommentScoreHideMins           int    `json:"comment_score_hide_mins"`
    AllowPredictions               bool   `json:"allow_predictions"`
    UserHasFavorited               bool   `json:"user_has_favorited"`
    UserFlairTemplateID            any    `json:"user_flair_template_id"`
    CommunityIcon                  string `json:"community_icon"`
    BannerBackgroundImage          string `json:"banner_background_image"`
    OriginalContentTagEnabled      bool   `json:"original_content_tag_enabled"`
    CommunityReviewed              bool   `json:"community_reviewed"`
    SubmitText                     string `json:"submit_text"`
    DescriptionHTML                string `json:"description_html"`
    SpoilersEnabled                bool   `json:"spoilers_enabled"`
    CommentContributionSettings    struct {
      AllowedMediaTypes []string `json:"allowed_media_types"`
    } `json:"comment_contribution_settings"`
    AllowTalks                       bool     `json:"allow_talks"`
    HeaderSize                       any      `json:"header_size"`
    UserFlairPosition                string   `json:"user_flair_position"`
    AllOriginalContent               bool     `json:"all_original_content"`
    HasMenuWidget                    bool     `json:"has_menu_widget"`
    IsEnrolledInNewModmail           any      `json:"is_enrolled_in_new_modmail"`
    KeyColor                         string   `json:"key_color"`
    CanAssignUserFlair               bool     `json:"can_assign_user_flair"`
    Created                          float64  `json:"created"`
    Wls                              int      `json:"wls"`
    ShowMediaPreview                 bool     `json:"show_media_preview"`
    SubmissionType                   string   `json:"submission_type"`
    UserIsSubscriber                 bool     `json:"user_is_subscriber"`
    AllowedMediaInComments           []string `json:"allowed_media_in_comments"`
    AllowVideogifs                   bool     `json:"allow_videogifs"`
    ShouldArchivePosts               bool     `json:"should_archive_posts"`
    UserFlairType                    string   `json:"user_flair_type"`
    AllowPolls                       bool     `json:"allow_polls"`
    CollapseDeletedComments          bool     `json:"collapse_deleted_comments"`
    EmojisCustomSize                 any      `json:"emojis_custom_size"`
    PublicDescriptionHTML            string   `json:"public_description_html"`
    AllowVideos                      bool     `json:"allow_videos"`
    IsCrosspostableSubreddit         any      `json:"is_crosspostable_subreddit"`
    NotificationLevel                string   `json:"notification_level"`
    ShouldShowMediaInCommentsSetting bool     `json:"should_show_media_in_comments_setting"`
    CanAssignLinkFlair               bool     `json:"can_assign_link_flair"`
    AccountsActiveIsFuzzed           bool     `json:"accounts_active_is_fuzzed"`
    AllowPredictionContributors      bool     `json:"allow_prediction_contributors"`
    SubmitTextLabel                  string   `json:"submit_text_label"`
    LinkFlairPosition                string   `json:"link_flair_position"`
    UserSrFlairEnabled               bool     `json:"user_sr_flair_enabled"`
    UserFlairEnabledInSr             bool     `json:"user_flair_enabled_in_sr"`
    AllowDiscovery                   bool     `json:"allow_discovery"`
    AcceptFollowers                  bool     `json:"accept_followers"`
    UserSrThemeEnabled               bool     `json:"user_sr_theme_enabled"`
    LinkFlairEnabled                 bool     `json:"link_flair_enabled"`
    DisableContributorRequests       bool     `json:"disable_contributor_requests"`
    SubredditType                    string   `json:"subreddit_type"`
    SuggestedCommentSort             any      `json:"suggested_comment_sort"`
    BannerImg                        string   `json:"banner_img"`
    UserFlairText                    any      `json:"user_flair_text"`
    BannerBackgroundColor            string   `json:"banner_background_color"`
    ShowMedia                        bool     `json:"show_media"`
    ID                               string   `json:"id"`
    UserIsModerator                  bool     `json:"user_is_moderator"`
    Over18                           bool     `json:"over18"`
    HeaderTitle                      string   `json:"header_title"`
    Description                      string   `json:"description"`
    SubmitLinkLabel                  string   `json:"submit_link_label"`
    UserFlairTextColor               any      `json:"user_flair_text_color"`
    RestrictCommenting               bool     `json:"restrict_commenting"`
    UserFlairCSSClass                any      `json:"user_flair_css_class"`
    AllowImages                      bool     `json:"allow_images"`
    Lang                             string   `json:"lang"`
    WhitelistStatus                  string   `json:"whitelist_status"`
    URL                              string   `json:"url"`
    CreatedUtc                       float64  `json:"created_utc"`
    BannerSize                       any      `json:"banner_size"`
    MobileBannerImage                string   `json:"mobile_banner_image"`
    UserIsContributor                bool     `json:"user_is_contributor"`
    AllowPredictionsTournament       bool     `json:"allow_predictions_tournament"`
  } `json:"data"`
}
