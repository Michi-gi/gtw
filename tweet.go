package gtw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type TweetsApi struct {
	OAuth1Client *AuthClient
	OAuth2Client *AuthClient
}

type expansionFields struct {
	Name   string
	Params []string
}

var mediaFields = expansionFields{
	Name:   "media.fields",
	Params: []string{"duration_ms", "height", "media_key", "preview_image_url", "type", "url", "width", "public_metrics", "non_public_metrics", "organic_metrics", "promoted_metrics", "alt_text"},
}

var placeFields = expansionFields{
	Name:   "place.fields",
	Params: []string{"contained_within", "country", "country_code", "full_name", "geo", "id", "name", "place_type"},
}

var pollFields = expansionFields{
	Name:   "poll.fields",
	Params: []string{"duration_minutes", "end_datetime", "id", "options", "voting_status"},
}

var tweetFields = expansionFields{
	Name:   "tweet.fields",
	Params: []string{"attachments", "author_id", "context_annotations", "conversation_id", "created_at", "entities", "geo", "id", "in_reply_to_user_id", "lang", "non_public_metrics", "public_metrics", "organic_metrics", "promoted_metrics", "possibly_sensitive", "referenced_tweets", "reply_settings", "source", "text", "withheld"},
}

var userFields = expansionFields{
	Name:   "user.fields",
	Params: []string{"created_at", "description", "entities", "id", "location", "name", "pinned_tweet_id", "profile_image_url", "protected", "public_metrics", "url", "username", "verified", "withheld"},
}

var expansionFieldsMap = map[string]expansionFields{
	"attachments.media_keys":         mediaFields,
	"geo.place_id":                   placeFields,
	"attachments.poll_ids":           pollFields,
	"referenced_tweets.id":           tweetFields,
	"author_id":                      userFields,
	"entities.mentions.username":     userFields,
	"in_reply_to_user_id":            userFields,
	"referenced_tweets.id.author_id": userFields,
}

func (t TweetsApi) selectClient(noUser bool) *AuthClient {
	if noUser || (t.OAuth1Client == nil) {
		return t.OAuth2Client
	} else {
		return t.OAuth1Client
	}
}

type TweetsLookupIdReturn struct {
	Data     Tweet `json:"data"`
	Includes struct {
		Tweets []Tweet `json:"tweets"`
		Users  []User  `json:"users"`
		Places []Place `json:"places"`
		Media  []Media `json:"media"`
		Polls  []Poll  `json:"polls"`
	} `json:"includes"`
	Errors Error `json:"errors"`
}

func (t TweetsApi) Get(id string, expansionParams []string, noUser bool) (*TweetsLookupIdReturn, error) {
	var url string
	if expansionParams == nil {
		url = fmt.Sprintf("https://api.twitter.com/2/tweets/%s", id)
	} else {
		expansionStr, expressions := mkFields(expansionParams)
		url = fmt.Sprintf("https://api.twitter.com/2/tweets/%s?expansions=%s&%s", id, expansionStr, strings.Join(expressions, "&"))
	}
	client := t.selectClient(noUser)
	req, err := client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Exec(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := new(TweetsLookupIdReturn)
	json.Unmarshal(bodyBytes, result)

	return result, nil
}

type TweetsLookupIdsReturn struct {
	Data     []Tweet `json:"data"`
	Includes struct {
		Tweets []Tweet `json:"tweets"`
		Users  []User  `json:"users"`
		Places []Place `json:"places"`
		Media  []Media `json:"media"`
		Polls  []Poll  `json:"polls"`
	} `json:"includes"`
	Errors Error `json:"errors"`
}

func (t TweetsApi) MGet(ids []string, expansionParams []string, noUser bool) (*TweetsLookupIdsReturn, error) {
	var url string
	concatedIds := strings.Join(ids, ",")
	if expansionParams == nil {
		url = fmt.Sprintf("https://api.twitter.com/2/tweets?ids=%s", concatedIds)
	} else {
		expansionStr, expressions := mkFields(expansionParams)
		url = fmt.Sprintf("https://api.twitter.com/2/tweets?ids=%s&expansions=%s&%s", concatedIds, expansionStr, strings.Join(expressions, "&"))
	}
	client := t.selectClient(noUser)
	req, err := client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Exec(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := new(TweetsLookupIdsReturn)
	json.Unmarshal(bodyBytes, result)

	return result, nil
}

type MediaPost struct {
	MediaIds      []string `json:"media_ids"`
	TaggedUserIds []string `json:"tagged_user_ids,omitempty"`
}
type PollPost struct {
	Options         []string `json:"options"`
	DurationMinutes int      `json:"duration_minutes"`
}
type ReplyPost struct {
	ExcludeReplyUserIds []string `json:"exclude_reply_user_ids,omitempty"`
	InReplyToTweetId    string   `json:"in_reply_to_tweet_id"`
}
type PostTweet struct {
	Text                  string     `json:"text"`
	Media                 *MediaPost `json:"media,omitempty"`
	Poll                  *PollPost  `json:"poll,omitempty"`
	QuoteTweetId          string     `json:"quoute_tweet_id,omitempty"`
	Reply                 *ReplyPost `json:"reply,omitempty"`
	DirectMessageDeepLink string     `json:"direct_message_deep_link,omitempty"`
	ReplySettings         string     `json:"reply_settings,omitempty"`
	ForSuperFollowersOnly bool       `json:"for_super_followers_only,omitempty"`
}

type postReturn struct {
	Data struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}

func (t TweetsApi) Post(postTweet PostTweet) (string, error) {
	bodyString, _ := json.Marshal(postTweet)
	req, err := t.OAuth1Client.NewRequest("POST", "https://api.twitter.com/2/tweets", strings.NewReader(string(bodyString)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := t.OAuth1Client.Exec(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	result := new(postReturn)
	json.Unmarshal(bodyBytes, result)

	return result.Data.Id, nil
}

type deleteReturn struct {
	Data struct {
		Deleted bool `json:"deleted"`
	} `json:"data"`
}

func (t TweetsApi) Delete(id string) (bool, error) {
	req, err := t.OAuth1Client.NewRequest("DELETE", fmt.Sprintf("https://api.twitter.com/2/tweets/%s", id), nil)
	if err != nil {
		return false, err
	}

	res, err := t.OAuth1Client.Exec(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	result := new(deleteReturn)
	json.Unmarshal(bodyBytes, result)

	return result.Data.Deleted, nil
}

type TimelineReturn struct {
	Data     []Tweet `json:"data"`
	Includes struct {
		Tweets []Tweet `json:"tweets"`
		Users  []User  `json:"users"`
		Places []Place `json:"places"`
		Media  []Media `json:"media"`
		Polls  []Poll  `json:"polls"`
	} `json:"includes"`
	Meta struct {
		Count         int    `json:"count"`
		NewestId      string `json:"newest_id"`
		OldestId      string `json:"oldest_id"`
		NextToken     string `json:"next_token"`
		PreviousToken string `json:"previous_token"`
	} `json:"meta"`
	Errors Error `json:"errors"`
}

func (t TweetsApi) Timelines(userId string, exclude []string, expansionParams []string, startTime *time.Time, endTime *time.Time, maxResults int, sinceId string, untilId string, pagenationToken string, noUser bool) (*TimelineReturn, error) {
	builder := urlBuilder{}
	builder.Url = fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets", userId)
	builder.appendParams("exclude", exclude)
	builder.appendParams("expansions", expansionParams)
	if startTime != nil {
		builder.appendParam("start_time", startTime.Format(time.RFC3339))
	}
	if endTime != nil {
		builder.appendParam("end_time", endTime.Format(time.RFC3339))
	}
	if maxResults != 0 {
		builder.appendParam("max_results", strconv.Itoa(maxResults))
	}
	builder.appendParam("since_id", sinceId)
	builder.appendParam("until_id", untilId)
	builder.appendParam("pagenation_token", pagenationToken)

	url := builder.Url
	if expansionParams != nil {
		_, expressions := mkFields(expansionParams)
		url = fmt.Sprintf(url+"&%s", strings.Join(expressions, "&"))
	}
	client := t.selectClient(noUser)
	req, err := client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Exec(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := new(TimelineReturn)
	json.Unmarshal(bodyBytes, result)

	return result, nil
}

func (t TweetsApi) Mentions(userId string, exclude []string, expansionParams []string, startTime *time.Time, endTime *time.Time, maxResults int, sinceId string, untilId string, pagenationToken string, noUser bool) (*TimelineReturn, error) {
	builder := urlBuilder{}
	builder.Url = fmt.Sprintf("https://api.twitter.com/2/users/%s/mentions", userId)
	builder.appendParams("exclude", exclude)
	builder.appendParams("expansions", expansionParams)
	if startTime != nil {
		builder.appendParam("start_time", startTime.Format(time.RFC3339))
	}
	if endTime != nil {
		builder.appendParam("end_time", endTime.Format(time.RFC3339))
	}
	if maxResults != 0 {
		builder.appendParam("max_results", strconv.Itoa(maxResults))
	}
	builder.appendParam("since_id", sinceId)
	builder.appendParam("until_id", untilId)
	builder.appendParam("pagenation_token", pagenationToken)

	url := builder.Url
	if expansionParams != nil {
		_, expressions := mkFields(expansionParams)
		url = fmt.Sprintf(url+"&%s", strings.Join(expressions, "&"))
	}
	client := t.selectClient(noUser)
	req, err := client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Exec(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := new(TimelineReturn)
	json.Unmarshal(bodyBytes, result)

	return result, nil
}

func (t TweetsApi) Search(query string, mode string, expansionParams []string, startTime *time.Time, endTime *time.Time, maxResults int, sinceId string, untilId string, pagenationToken string, noUser bool) (*TimelineReturn, error) {
	builder := urlBuilder{}
	builder.Url = fmt.Sprintf("https://api.twitter.com/2/tweets/search/%s", mode)
	builder.appendParam("query", query)
	builder.appendParams("expansions", expansionParams)
	if startTime != nil {
		builder.appendParam("start_time", startTime.Format(time.RFC3339))
	}
	if endTime != nil {
		builder.appendParam("end_time", endTime.Format(time.RFC3339))
	}
	if maxResults != 0 {
		builder.appendParam("max_results", strconv.Itoa(maxResults))
	}
	builder.appendParam("since_id", sinceId)
	builder.appendParam("until_id", untilId)
	builder.appendParam("pagenation_token", pagenationToken)

	url := builder.Url
	if expansionParams != nil {
		_, expressions := mkFields(expansionParams)
		url = fmt.Sprintf(url+"&%s", strings.Join(expressions, "&"))
	}
	client := t.selectClient(noUser)
	req, err := client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Exec(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := new(TimelineReturn)
	json.Unmarshal(bodyBytes, result)

	return result, nil
}

func mkFields(fieldParams []string) (string, []string) {
	var result []string
	for _, field := range fieldParams {
		result = append(result, expansionFieldsMap[field].Name+"="+strings.Join(expansionFieldsMap[field].Params, ","))
	}

	return strings.Join(fieldParams, ","), result
}

type urlBuilder struct {
	Url string
}

func (b *urlBuilder) appendParam(key string, param string) {
	if param != "" {
		if strings.Contains(b.Url, "?") {
			b.Url += "&"
		} else {
			b.Url += "?"
		}
		b.Url += (key + "=" + param)
	}
}

func (b *urlBuilder) appendParams(key string, params []string) {
	if len(params) > 0 {
		if strings.Contains(b.Url, "?") {
			b.Url += "&"
		} else {
			b.Url += "?"
		}
		b.Url += (key + "=" + strings.Join(params, ","))
	}
}
