package main

// imports ...

type TestCommand struct{}

var _ ken.Command = (*TestCommand)(nil)

func (c *TestCommand) Name() string {
	return "ping"
}

func (c *TestCommand) Description() string {
	return "Basic Ping Command"
}

func (c *TestCommand) Version() string {
	return "1.0.0"
}

func (c *TestCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *TestCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *TestCommand) Run(ctx *ken.Ctx) (err error) {
	err = ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
	return
}

func main() {
	token := os.Getenv("TOKEN")

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	k := ken.New(session)
	k.RegisterCommands(
		new(commands.TestCommand),
	)

	defer k.Unregister()

	err = session.Open()
	if err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
