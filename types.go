package main

type Method int

const (
	TorrentGet Method = iota
	TorrentAdd
	TorrentRenamePath
)

func (c Method) String() string {
	return [...]string{
		"torrent-get",
		"torrent-add",
		"torrent-rename-path",
	}[c]
}

type Peer struct {
	Address            string  `json:"address,omitempty"`
	ClientName         string  `json:"clientName,omitempty"`
	ClientIsChoked     bool    `json:"clientIsChoked,omitempty"`
	ClientIsInterested bool    `json:"clientIsInterested,omitempty"`
	FlagString         string  `json:"flagStr,omitempty"`
	IsDownloadingFrom  bool    `json:"isDownloadingFrom,omitempty"`
	IsEncrypted        bool    `json:"isEncrypted,omitempty"`
	IsIncoming         bool    `json:"isIncoming,omitempty"`
	IsUploadingTo      bool    `json:"isUploadingTo,omitempty"`
	IsUTP              bool    `json:"isUTP,omitempty"`
	PeerIsChoked       bool    `json:"peerIsChoked,omitempty"`
	PeerIsInterested   bool    `json:"peerIsInterested,omitempty"`
	Port               int     `json:"port,omitempty"`
	Progress           float32 `json:"progress,omitempty"`
	RateToClient       int     `json:"rateToClient,omitempty"`
	RateToPeer         int     `json:"rateToPeer,omitempty"`
}

type File struct {
	BytesCompleted int    `json:"bytesCompleted,omitempty"`
	Length         int    `json:"length,omitempty"`
	Name           string `json:"name,omitempty"`
}

type FileStat struct {
	BytesCompleted int  `json:"bytesCompleted,omitempty"`
	Wanted         bool `json:"wanted,omitempty"`
	Priority       int  `json:"priority,omitempty"`
}

type PeersFrom struct {
	FromCache    uint `json:"fromCache,omitempty"`
	FromDHT      uint `json:"fromDht,omitempty"`
	FromIncoming uint `json:"fromIncoming,omitempty"`
	FromLPD      uint `json:"fromLpd,omitempty"`
	FromLTEP     uint `json:"fromLtep,omitempty"`
	FromPEX      uint `json:"fromPex,omitempty"`
	FromTracker  uint `json:"fromTracker,omitempty"`
}

type Tracker struct {
	Announce string `json:"announce,omitempty"`
	ID       int    `json:"id,omitempty"`
	Scrape   string `json:"scrape,omitempty"`
	Tier     int    `json:"tier,omitempty"`
}

type TrackerStat struct {
	Announce              string `json:"announce,omitempty"`
	AnnounceState         int    `json:"announceState,omitempty"`
	DownloadCount         int    `json:"downloadCount,omitempty"`
	HasAnnounced          bool   `json:"hasAnnounced,omitempty"`
	HasScraped            bool   `json:"hasScraped,omitempty"`
	Host                  string `json:"host,omitempty"`
	ID                    int    `json:"id,omitempty"`
	IsBackup              bool   `json:"isBackup,omitempty"`
	LastAnnouncePeerCount int    `json:"lastAnnouncePeerCount,omitempty"`
	LastAnnounceResult    string `json:"lastAnnounceResult,omitempty"`
	LastAnnounceStartTime int    `json:"lastAnnounceStartTime,omitempty"`
	LastAnnounceSucceeded bool   `json:"lastAnnounceSucceeded,omitempty"`
	LastAnnounceTime      int    `json:"lastAnnounceTime,omitempty"`
	LastAnnounceTimedOut  bool   `json:"lastAnnounceTimedOut,omitempty"`
	LastScrapeResult      string `json:"lastScrapeResult,omitempty"`
	LastScrapeStartTime   int    `json:"lastScrapeStartTime,omitempty"`
	LastScrapeSucceeded   bool   `json:"lastScrapeSucceeded,omitempty"`
	LastScrapeTime        int    `json:"lastScrapeTime,omitempty"`
	LastScrapeTimedOut    bool   `json:"lastScrapeTimedOut,omitempty"`
	LeecherCount          int    `json:"leecherCount,omitempty"`
	NextAnnounceTime      int    `json:"nextAnnounceTime,omitempty"`
	NextScrapeTime        int    `json:"nextScrapeTime,omitempty"`
	Scrape                string `json:"scrape,omitempty"`
	ScrapeState           int    `json:"scrapeState,omitempty"`
	SeederCount           int    `json:"seederCount,omitempty"`
	Tier                  int    `json:"tier,omitempty"`
}

type Torrent struct {
	ActivityDate            int           `json:"activityDate,omitempty"`
	AddedDate               int           `json:"addedDate,omitempty"`
	BandwidthPriority       int           `json:"bandwidthPriority,omitempty"`
	Comment                 string        `json:"comment,omitempty"`
	CorruptEver             int           `json:"corruptEver,omitempty"`
	Creator                 string        `json:"creator,omitempty"`
	DateCreated             int           `json:"dateCreated,omitempty"`
	DesiredAvailable        int           `json:"desiredAvailable,omitempty"`
	DoneDate                int           `json:"doneDate,omitempty"`
	DownloadDir             string        `json:"downloadDir,omitempty"`
	DownloadedEver          int           `json:"downloadedEver,omitempty"`
	DownloadLimit           int           `json:"downloadLimit,omitempty"`
	DownloadLimited         bool          `json:"downloadLimited,omitempty"`
	Error                   int           `json:"error,omitempty"`
	ErrorString             string        `json:"errorString,omitempty"`
	ETA                     int           `json:"eta,omitempty"`
	ETAIdle                 int           `json:"etaIdle,omitempty"`
	Files                   []File        `json:"files,omitempty"`
	FileStats               []FileStat    `json:"fileStats,omitempty"`
	HashString              string        `json:"hashString,omitempty"`
	HaveUnchecked           int           `json:"haveUnchecked,omitempty"`
	HaveValid               int           `json:"haveValid,omitempty"`
	HonorsSessionLimits     bool          `json:"honorsSessionLimits,omitempty"`
	ID                      int           `json:"id,omitempty"`
	IsFinished              bool          `json:"isFinished,omitempty"`
	IsPrivate               bool          `json:"isPrivate,omitempty"`
	IsStalled               bool          `json:"isStalled,omitempty"`
	Labels                  []string      `json:"labels,omitempty"`
	LeftUntilDone           int           `json:"leftUntilDone,omitempty"`
	MagnetLink              string        `json:"magnetLink,omitempty"`
	ManualAnnounceTime      int           `json:"manualAnnounceTime,omitempty"`
	MaxConnectedPeers       int           `json:"maxConnectedPeers,omitempty"`
	MetadataPercentComplete float64       `json:"metadataPercentComplete,omitempty"`
	Name                    string        `json:"name,omitempty"`
	Peer                    int           `json:"peer,omitempty"`
	Peers                   []Peer        `json:"peers,omitempty"`
	PeersConnected          int           `json:"peersConnected,omitempty"`
	PeersFrom               PeersFrom     `json:"peersFrom,omitempty"`
	PeersGettingFromUs      int           `json:"peersGettingFromUs,omitempty"`
	PeersSendingToUs        int           `json:"peersSendingToUs,omitempty"`
	PercentDone             float64       `json:"percentDone,omitempty"`
	Pieces                  string        `json:"pieces,omitempty"`
	PieceCount              int           `json:"pieceCount,omitempty"`
	PieceSize               int           `json:"pieceSize,omitempty"`
	Priorities              []uint        `json:"priorities,omitempty"`
	QueuePosition           int           `json:"queuePosition,omitempty"`
	RateDownload            int           `json:"rateDownload,omitempty"`
	RateUpload              int           `json:"rateUpload,omitempty"`
	RecheckProgress         float64       `json:"recheckProgress,omitempty"`
	SecondsDownloading      int           `json:"secondsDownloading,omitempty"`
	SecondsSeeding          int           `json:"secondsSeeding,omitempty"`
	SeedIdleLimit           int           `json:"seedIdleLimit,omitempty"`
	SeedIdleMode            int           `json:"seedIdleMode,omitempty"`
	SeedRatioLimit          float64       `json:"seedRatioLimit,omitempty"`
	SeedRatioMode           int           `json:"seedRatioMode,omitempty"`
	SizeWhenDone            int           `json:"sizeWhenDone,omitempty"`
	StartDate               int           `json:"startDate,omitempty"`
	Status                  int           `json:"status,omitempty"`
	Trackers                []Tracker     `json:"trackers,omitempty"`
	TrackerStats            []TrackerStat `json:"trackerStats,omitempty"`
	TotalSize               int           `json:"totalSize,omitempty"`
	TorrentFile             string        `json:"torrentFile,omitempty"`
	UploadedEver            int           `json:"uploadedEver,omitempty"`
	UploadLimit             int           `json:"uploadLimit,omitempty"`
	UploadLimited           bool          `json:"uploadLimited,omitempty"`
	UploadRatio             float64       `json:"uploadRatio,omitempty"`
	Wanted                  []bool        `json:"wanted,omitempty"`
	Webseeds                []string      `json:"webseeds,omitempty"`
	WebseedsSendingToUs     int           `json:"webseedsSendingToUs,omitempty"`
}

type TorrentGetRequest struct {
	Method    string `json:"method"`
	Arguments struct {
		IDs    []int    `json:"ids,omitempty"`
		Fields []string `json:"fields"`
	} `json:"arguments"`
}

type TorrentGetResponse struct {
	Arguments struct {
		Torrents []Torrent `json:"torrents"`
	} `json:"arguments"`
}

type TorrentAddRequest struct {
	Method    string `json:"method"`
	Arguments struct {
		Cookies           string   `json:"cookies,omitempty"`
		DownloadDir       string   `json:"download-dir,omitempty"`
		Filename          string   `json:"filename,omitempty"`
		Metainfo          string   `json:"metainfo,omitempty"`
		Paused            bool     `json:"paused,omitempty"`
		PeerLimit         uint     `json:"peer-limit,omitempty"`
		BandwidthPriority uint     `json:"bandwidthPriority,omitempty"`
		FilesWanted       []string `json:"files-wanted,omitempty"`
		FilesUnwanted     []string `json:"files-wanted,omitempty"`
		PriorityHigh      []string `json:"priority-high,omitempty"`
		PriorityLow       []string `json:"priority-low,omitempty"`
		PriorityNormal    []string `json:"priority-normal,omitempty"`
	} `json:"arguments"`
}

type TorrentAddResponse struct {
	Method    string `json:"method"`
	Arguments struct {
		HashString string `json:"hashString,omitempty"`
		ID         int    `json:"id,omitempty"`
		Name       string `json:"name,omitempty"`
	} `json:"arguments"`
}

type TorrentRenamePathRequest struct {
	Method    string `json:"method"`
	Arguments struct {
		ID   int    `json:"ids"`
		Path string `json:"path"`
		Name string `json:"name"`
	} `json:"arguments"`
}

type TorrentRenamePathResponse struct {
	Method    string `json:"method"`
	Arguments struct {
		ID   int    `json:"id"`
		Path string `json:"path"`
		Name string `json:"name"`
	} `json:"arguments"`
}
