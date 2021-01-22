package cli

import (
	"os"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	KafkaDefaultConnectTimeout   = "10s"
	KafkaDefaultGroupId          = "plumber"
	KafkaDefaultMaxWait          = "1s"
	KafkaDefaultMinBytes         = "1"
	KafkaDefaultMaxBytes         = "1"
	KafkaDefaultQueueCapacity    = "1"
	KafkaDefaultRebalanceTimeout = "0"
)

type KafkaOptions struct {
	// Shared
	Address            string
	Topic              string
	Timeout            time.Duration
	InsecureTLS        bool
	Username           string
	Password           string
	AuthenticationType string

	// Read
	ReadGroupId      string
	MaxWait          time.Duration
	MinBytes         int
	MaxBytes         int
	QueueCapacity    int
	RebalanceTimeout time.Duration

	// Write
	WriteKey string
}

func HandleKafkaFlags(readCmd, writeCmd, relayCmd *kingpin.CmdClause, opts *Options) {
	rc := readCmd.Command("kafka", "Kafka message system")

	addSharedKafkaFlags(rc, opts)
	addReadKafkaFlags(rc, opts)

	// Kafka write cmd
	wc := writeCmd.Command("kafka", "Kafka message system")

	addSharedKafkaFlags(wc, opts)
	addWriteKafkaFlags(wc, opts)

	// If PLUMBER_RELAY_TYPE is set, use env vars, otherwise use CLI flags
	relayType := os.Getenv("PLUMBER_RELAY_TYPE")

	var rec *kingpin.CmdClause

	if relayType != "" {
		rec = relayCmd
	} else {
		rec = relayCmd.Command("kafka", "Kafka message system")
	}

	addSharedKafkaFlags(rec, opts)
	addReadKafkaFlags(rec, opts)
}

func addSharedKafkaFlags(cmd *kingpin.CmdClause, opts *Options) {
	cmd.Flag("address", "Destination host address").
		Default("localhost:9092").
		Envar("PLUMBER_RELAY_KAFKA_ADDRESS").
		StringVar(&opts.Kafka.Address)
	cmd.Flag("topic", "Topic to read message(s) from").
		Required().
		Envar("PLUMBER_RELAY_KAFKA_TOPIC").
		StringVar(&opts.Kafka.Topic)
	cmd.Flag("timeout", "Connect timeout").Default(KafkaDefaultConnectTimeout).
		Envar("PLUMBER_RELAY_KAFKA_TIMEOUT").
		DurationVar(&opts.Kafka.Timeout)
	cmd.Flag("insecure-tls", "Use insecure TLS").
		Envar("PLUMBER_RELAY_KAFKA_INSECURE_TLS").
		BoolVar(&opts.Kafka.InsecureTLS)
	cmd.Flag("username", "SASL Username").
		Envar("PLUMBER_RELAY_KAFKA_USERNAME").
		StringVar(&opts.Kafka.Username)
	cmd.Flag("password", "SASL Password. If omitted, you will be prompted for the password").
		Envar("PLUMBER_RELAY_KAFKA_PASSWORD").
		StringVar(&opts.Kafka.Password)
	cmd.Flag("auth-type", "SASL Authentication type (plain or scram)").
		Default("scram").
		Envar("PLUMBER_RELAY_KAFKA_SASL_TYPE").
		StringVar(&opts.Kafka.AuthenticationType)
}

func addReadKafkaFlags(cmd *kingpin.CmdClause, opts *Options) {
	cmd.Flag("group-id", "Specify a specific group-id to use when reading from kafka").
		Envar("PLUMBER_RELAY_GROUP_ID").
		Default(KafkaDefaultGroupId).StringVar(&opts.Kafka.ReadGroupId)
	cmd.Flag("max-wait", "How long to wait for new data when reading batches of messages").
		Envar("PLUMBER_RELAY_MAX_WAIT").
		Default(KafkaDefaultMaxWait).DurationVar(&opts.Kafka.MaxWait)
	cmd.Flag("min-bytes", "Minimum number of bytes to fetch in a single kafka request (throughput optimization)").
		Envar("PLUMBER_RELAY_MIN_BYTES").
		Default(KafkaDefaultMinBytes).IntVar(&opts.Kafka.MinBytes)
	cmd.Flag("max-bytes", "Maximum number of bytes to fetch in a single kafka request (throughput optimization)").
		Envar("PLUMBER_RELAY_MAX_BYTES").
		Default(KafkaDefaultMaxBytes).IntVar(&opts.Kafka.MaxBytes)
	cmd.Flag("queue-capacity", "Internal queue capacity").
		Envar("PLUMBER_RELAY_QUEUE_CAPACITY").
		Default(KafkaDefaultQueueCapacity).IntVar(&opts.Kafka.QueueCapacity)
	cmd.Flag("rebalance-timeout", "How long a coordinator will wait for member joins as part of a rebalance").
		Envar("PLUMBER_RELAY_REBALANCE_TIMEOUT").
		Default(KafkaDefaultRebalanceTimeout).DurationVar(&opts.Kafka.RebalanceTimeout)
}

func addWriteKafkaFlags(cmd *kingpin.CmdClause, opts *Options) {
	cmd.Flag("key", "Key to write to kafka (not required)").StringVar(&opts.Kafka.WriteKey)
}
