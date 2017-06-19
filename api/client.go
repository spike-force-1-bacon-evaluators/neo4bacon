package api

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/spf13/viper"
)

// Client wraps neo4j connection values
type Client struct {
	user      string
	password  string
	host      string
	httpPort  string
	boltPort  string
	url       string
	Conn      bolt.Conn
	CloseConn func() error
}

// NewClient returns a Neo4j Client
func NewClient(filepath string) (Client, error) {
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		return Client{}, fmt.Errorf("error loading config file: %s", err)
	}

	// Load neo4j config values
	neo4j, err := loadNeo4j()
	if err != nil {
		return Client{}, fmt.Errorf("failed to load configurations: %s", err)
	}

	// We don't need Viper anymore since the config values are already loaded
	viper.Reset()
	return *neo4j, nil
}

// loadNeo4j connection values to through Viper
func loadNeo4j() (*Client, error) {
	keys := []string{
		"user",
		"password",
		"host",
		"http_port",
		"bolt_port",
	}

	// Check if values are available
	for _, v := range keys {
		if !viper.IsSet(fmt.Sprintf("neo4j.%s", v)) {
			return nil, fmt.Errorf("could not load: %s", v)
		}
	}

	return &Client{
		user:     viper.GetString("neo4j.user"),
		password: viper.GetString("neo4j.password"),
		host:     viper.GetString("neo4j.host"),
		httpPort: viper.GetString("neo4j.http_port"),
		boltPort: viper.GetString("neo4j.bolt_port"),
	}, nil
}

// Returns the connection url
func (n *Client) connURL() {
	n.url = fmt.Sprintf("bolt://%s:%s@%s:%s/db/data",
		n.user, n.password, n.host, n.boltPort)
}

// Connection returns a new connection to Neo4j, a close function
// to interrupt the connection and an error
func (n *Client) Connection() error {
	n.connURL()
	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo(n.url)
	if err != nil {
		return fmt.Errorf("failed to open neo4j connection: %s", err)
	}

	n.Conn = conn
	n.CloseConn = conn.Close
	return nil
}
