package gtw

import (
	"time"
)

type Tweet struct {
	Id          string `json:"id"`
	Text        string `json:"text"`
	Attachments struct {
		MediaKeys []string `json:"media_keys"`
		PollIds   []string `json:"poll_ids"`
	} `json:"attachments"`
	AuthorId           string `json:"author_id"`
	ContextAnnotations []struct {
		Domain struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"domain"`
		Entity struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"entity"`
	} `json:"content_annotations"`
	ConversationId string    `json:"conversation_id"`
	CreatedAt      time.Time `json:"created_at"`
	Entities       struct {
		Annotations []struct {
			Start          int     `json:"start"`
			End            int     `json:"end"`
			Probability    float32 `json:"probability"`
			Type           string  `json:"type"`
			NormalizedText string  `json:"normalized_text"`
		} `json:"annotations"`
		Urls []struct {
			Start       int    `json:"start"`
			End         int    `json:"end"`
			Url         string `json:"url"`
			Expand      string `json:"expand"`
			DisplayUrl  string `json:"display_url"`
			UnknowndUrl string `json:"unknown_url"`
		} `json:"urls"`
		Hashtags []struct {
			Start int    `json:"start"`
			End   int    `json:"end"`
			Tag   string `json:"tag"`
		} `json:"hashtags"`
		Mentions []struct {
			Start    int    `json:"start"`
			End      int    `json:"end"`
			Username string `json:"username"`
		} `json:"mentions"`
		Cashtags []struct {
			Start int    `json:"start"`
			End   int    `json:"end"`
			Tag   string `json:"tag"`
		} `json:"cashtags"`
	} `json:"entities"`
	Geo struct {
		Coordinates struct {
			Type        string    `json:"type"`
			Coordinates []float32 `json:"coordinates"`
		} `json:"coordinates"`
		PlaceId string `json:"place_id"`
	} `json:"geo"`
	InReplyToUserId  string `json:"in_reply_to_user_id"`
	Lang             string `json:"lang"`
	NonPublicMetrics struct {
		ImpressionCount  int `json:"impression_count"`
		UrlLinkClicks    int `json:"url_link_clicks"`
		UserProfileClick int `json:"user_profile_click"`
	} `json:"non_public_metrics"`
	OrganicMetrics struct {
		ImpressionCount  int `json:"impression_count"`
		LikeCount        int `json:"like_count"`
		ReplyCount       int `json:"reply_count"`
		RetweetCount     int `json:"retweet_count"`
		UrlLinkClicks    int `json:"url_link_clicks"`
		UserProfileClick int `json:"user_profile_click"`
	} `json:"organic_metrics"`
	PossiblySensitive bool `json:"possibly_sensitive"`
	PromotedMetrics   struct {
		ImpressionCount  int `json:"impression_count"`
		LikeCount        int `json:"like_count"`
		ReplyCount       int `json:"reply_count"`
		RetweetCount     int `json:"retweet_count"`
		UrlLinkClicks    int `json:"url_link_clicks"`
		UserProfileClick int `json:"user_profile_click"`
	} `json:"promoted_metrics"`
	PublicMetrics struct {
		RetweetCount int `json:"retweet_count"`
		ReplyCount   int `json:"reply_count"`
		LikeCount    int `json:"like_count"`
		QuoteCount   int `json:"quote_count"`
	} `json:"public_metrics"`
	RefererncedTweets []struct {
		Type string `json:"type"`
		Id   string `json:"id"`
	} `json:"referenced_tweets"`
	ReplySettings string `json:"reply_settings"`
	Source        string `json:"source"`
	Withheld      struct {
		Copyright    bool     `json:"copyright"`
		CountryCodes []string `json:"country_codes"`
		Scope        string   `json:"scope"`
	} `json:"withheld"`
}

type User struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Entities    struct {
		Url []struct {
			Urls []struct {
				Start       int    `json:"start"`
				End         int    `json:"end"`
				Url         string `json:"url"`
				ExpandedUrl string `json:"expanded_url"`
				DisplayUrl  string `json:"display_url"`
			} `json:"urls"`
		} `json:"url"`
		Description []struct {
			Urls []struct {
				Start       int    `json:"start"`
				End         int    `json:"end"`
				Url         string `json:"url"`
				ExpandedUrl string `json:"expanded_url"`
				DisplayUrl  string `json:"display_url"`
			} `json:"urls"`
			Hashtags []struct {
				Start   int    `json:"start"`
				End     int    `json:"end"`
				Hashtag string `json:"hashtag"`
			} `json:"hashtags"`
			Mentions []struct {
				Start    int    `json:"start"`
				End      int    `json:"end"`
				Username string `json:"username"`
			} `json:"mentions"`
			Cashtags []struct {
				Start   int    `json:"start"`
				End     int    `json:"end"`
				Cashtag string `json:"cashtag"`
			} `json:"cashtags"`
		} `json:"description"`
	} `json:"entities"`
	Location        string `json:"location"`
	PinnedTweetId   string `json:"pinned_tweet_id"`
	ProfileImageUrl string `json:"profile_image_url"`
	Protected       bool   `json:"protected"`
	PublicMetrics   struct {
		FollowersCount int `json:"followers_count"`
		FollowingCount int `json:"following_count"`
		TweetCount     int `json:"tweet_count"`
		ListedCount    int `json:"listed_count"`
	} `json:"public_metrics"`
	Url      string `json:"url"`
	Verified bool   `json:"verified"`
	Withheld struct {
		CountryCodes []string `json:"country_codes"`
		Scope        string   `json:"scope"`
	} `json:"withheld"`
}

type Space struct {
	Id               string    `json:"id"`
	State            string    `json:"state"`
	CreatedAt        time.Time `json:"created_at"`
	EndedAt          time.Time `json:"ended_at"`
	HostIds          []string  `json:"host_ids"`
	Lang             string    `json:"lang"`
	IsTicketed       bool      `json:"is_ticketed"`
	InvitedUserIds   []string  `json:"invited_user_ids"`
	ParticipantCount int       `json:"participant_count"`
	ScheduledStart   time.Time `json:"scheduled_start"`
	SpeakerIds       []string  `json:"speaker_ids"`
	StartedAt        time.Time `json:"started_at"`
	Title            string    `json:"title"`
	TopicIds         []string  `json:"topic_ids"`
	CreatorId        string    `json:"creator_id"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type List struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"created_at"`
	Description   string    `json:"description"`
	FollowerCount int       `json:"follower_count"`
	MemberCount   int       `json:"member_count"`
	Private       bool      `json:"private"`
	OwnerId       string    `json:"owner_id"`
}

type Media struct {
	MediaKey         string `json:"media_key"`
	Type             string `json:"type"`
	DurationMs       int    `json:"duration_ms"`
	Height           int    `json:"height"`
	NonPublicMetrics struct {
		Playback0Count   int `json:"playback_0_count"`
		Playback25Count  int `json:"playback_25_count"`
		Playback50Count  int `json:"playback_50_count"`
		Playback75Count  int `json:"playback_75_count"`
		Playback100Count int `json:"playback_100_count"`
	} `json:"non_public_metrics"`
	OrganicMetrics struct {
		Playback0Count   int `json:"playback_0_count"`
		Playback25Count  int `json:"playback_25_count"`
		Playback50Count  int `json:"playback_50_count"`
		Playback75Count  int `json:"playback_75_count"`
		Playback100Count int `json:"playback_100_count"`
		ViewCount        int `json:"view_count"`
	} `json:"organic_metrics"`
	PreviewImageUrl string `json:"preview_image_url"`
	PromotedMetrics struct {
		Playback0Count   int `json:"playback_0_count"`
		Playback25Count  int `json:"playback_25_count"`
		Playback50Count  int `json:"playback_50_count"`
		Playback75Count  int `json:"playback_75_count"`
		Playback100Count int `json:"playback_100_count"`
		ViewCount        int `json:"view_count"`
	} `json:"promoted_metrics"`
	PublicMetrics struct {
		ViewCount int `json:"view_count"`
	} `json:"public_metrics"`
	Width   int    `json:"width"`
	AltText string `json:"alt_text"`
}

type Poll struct {
	Id      string `json:"id"`
	Options []struct {
		Position int    `json:"position"`
		Label    string `json:"label"`
		Votes    int    `json:"votes"`
	} `json:"options"`
	DurationMinutes int       `json:"duration_minutes"`
	EndDatetime     time.Time `json:"end_datetime"`
	VotingStatus    string    `json:"voting_status"`
}

type Place struct {
	FullName        string   `json:"full_name"`
	Id              string   `json:"id"`
	ContainedWithin []string `json:"contained_within"`
	Country         string   `json:"country"`
	CountryCode     string   `json:"country_code"`
	Geo             struct {
		Type       string    `json:"type"`
		Bbox       []float32 `json:"bbox"`
		Properties struct{}  `json:"properties"`
	} `json:"geo"`
	Name      string `json:"name"`
	PlaceType string `json:"place_type"`
}

type Error struct {
	Errors []struct {
		Parameters struct {
			EndTime []time.Time `json:"end_time"`
		} `json:"parameters"`
		Message string `json:"message"`
	} `json:"errors"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Type   string `json:"type"`
}
