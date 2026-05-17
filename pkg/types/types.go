package types

type Record struct {
	ID          int
	ArtistID    int
	Title       string
	ReleaseDate string
	Country     string
	Barcode     string
}

type Item struct {
	ID          int    `json:"id"`
	Status      string `json:"status"`
	Year        int    `json:"year"`
	ResourceURL string `json:"resource_url"`
	URI         string `json:"uri"`
	Artists     []struct {
		Name string `json:"name"`
		Anv  string `json:"anv"`
		Join string `json:"join"`
		Role string `json:"role"`
		// Tracks       string `json:"tracks"`
		ID           int    `json:"id"`
		ResourceURL  string `json:"resource_url"`
		ThumbnailURL string `json:"thumbnail_url"`
	} `json:"artists"`
	ArtistsSort string `json:"artists_sort"`
	Labels      []struct {
		Name           string `json:"name"`
		Catno          string `json:"catno"`
		EntityType     string `json:"entity_type"`
		EntityTypeName string `json:"entity_type_name"`
		ID             int    `json:"id"`
		ResourceURL    string `json:"resource_url"`
		ThumbnailURL   string `json:"thumbnail_url"`
	} `json:"labels"`
	Series    []any `json:"series"`
	Companies []struct {
		Name           string `json:"name"`
		Catno          string `json:"catno"`
		EntityType     string `json:"entity_type"`
		EntityTypeName string `json:"entity_type_name"`
		ID             int    `json:"id"`
		ResourceURL    string `json:"resource_url"`
		ThumbnailURL   string `json:"thumbnail_url,omitempty"`
	} `json:"companies"`
	Formats []struct {
		Name         string   `json:"name"`
		Qty          string   `json:"qty"`
		Descriptions []string `json:"descriptions"`
	} `json:"formats"`
	DataQuality       string  `json:"data_quality"`
	FormatQuantity    int     `json:"format_quantity"`
	DateAdded         string  `json:"date_added"`
	DateChanged       string  `json:"date_changed"`
	NumForSale        int     `json:"num_for_sale"`
	LowestPrice       float64 `json:"lowest_price"`
	MasterID          int     `json:"master_id"`
	MasterURL         string  `json:"master_url"`
	Title             string  `json:"title"`
	Country           string  `json:"country"`
	Released          string  `json:"released"`
	Notes             string  `json:"notes"`
	ReleasedFormatted string  `json:"released_formatted"`
	Identifiers       []struct {
		Type        string `json:"type"`
		Value       string `json:"value"`
		Description string `json:"description,omitempty"`
	} `json:"identifiers"`
	Genres    []string `json:"genres"`
	Styles    []string `json:"styles"`
	Tracklist []struct {
		Position string `json:"position"`
		Type     string `json:"type_"`
		Title    string `json:"title"`
		Duration string `json:"duration"`
	} `json:"tracklist"`
	// Extraartists []struct {
	// 	Name        string `json:"name"`
	// 	Anv         string `json:"anv"`
	// 	Join        string `json:"join"`
	// 	Role        string `json:"role"`
	// 	Tracks      string `json:"tracks"`
	// 	ID          int    `json:"id"`
	// 	ResourceURL string `json:"resource_url"`
	// } `json:"extraartists"`
	Thumb           string `json:"thumb"`
	EstimatedWeight int    `json:"estimated_weight"`
	BlockedFromSale bool   `json:"blocked_from_sale"`
	IsOffensive     bool   `json:"is_offensive"`
}

type Folder struct {
	// Pagination struct {
	// Page    int      `json:"page"`
	// Pages   int      `json:"pages"`
	// PerPage int      `json:"per_page"`
	// Items   int      `json:"items"`
	// Urls    struct{} `json:"urls"`
	// } `json:"pagination"`
	Releases []struct {
		ID               int    `json:"id"`
		InstanceID       int    `json:"instance_id"`
		DateAdded        string `json:"date_added"`
		Rating           int    `json:"rating"`
		BasicInformation struct {
			ID          int    `json:"id"`
			MasterID    int    `json:"master_id"`
			MasterURL   string `json:"master_url"`
			ResourceURL string `json:"resource_url"`
			// Thumb       string `json:"thumb"`
			// CoverImage  string `json:"cover_image"`
			Title   string `json:"title"`
			Year    int    `json:"year"`
			Formats []struct {
				Name         string   `json:"name"`
				Qty          string   `json:"qty"`
				Text         string   `json:"text"`
				Descriptions []string `json:"descriptions"`
			} `json:"formats"`
			Labels []struct {
				Name           string `json:"name"`
				Catno          string `json:"catno"`
				EntityType     string `json:"entity_type"`
				EntityTypeName string `json:"entity_type_name"`
				ID             int    `json:"id"`
				// ResourceURL    string `json:"resource_url"`
			} `json:"labels"`
			Artists []struct {
				Name   string `json:"name"`
				Anv    string `json:"anv"`
				Join   string `json:"join"`
				Role   string `json:"role"`
				Tracks string `json:"tracks"`
				ID     int    `json:"id"`
				// ResourceURL string `json:"resource_url"`
			} `json:"artists"`
			Genres []string `json:"genres"`
			Styles []any    `json:"styles"`
		} `json:"basic_information"`
		FolderID int `json:"folder_id"`
		Notes    []struct {
			FieldID int    `json:"field_id"`
			Value   string `json:"value"`
		} `json:"notes,omitempty"`
	} `json:"releases"`
}

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	ResourceURL  string `json:"resource_url"`
	ConsumerName string `json:"consumer_name"`
}
