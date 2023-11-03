package keyword

import (
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
    "log"
)

func (c *OzonKeywordCommander) Help(inputMessage *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
        "/help - help\n"+
            "/list - list products",
    )

    _, err := c.bot.Send(msg)
    if err != nil {
        log.Printf("OzonKeywordCommander.Help: error sending reply message to chat - %v", err)
    }
}
