package ozon

import (
    "github.com/ozonmp/omp-bot/internal/app/commands/ozon/keyword"
    "log"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
    HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
    HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type OzonCommander struct {
    bot              *tgbotapi.BotAPI
    keywordCommander Commander
}

func NewOzonCommander(
    bot *tgbotapi.BotAPI,
) *OzonCommander {
    return &OzonCommander{
        bot: bot,
        // keywordCommander
        keywordCommander: keyword.NewOzonKeywordCommander(bot),
    }
}

func (c *OzonCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
    switch callbackPath.Subdomain {
    case "keyword":
        c.keywordCommander.HandleCallback(callback, callbackPath)
    default:
        log.Printf("OzonCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
    }
}

func (c *OzonCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
    switch commandPath.Subdomain {
    case "keyword":
        c.keywordCommander.HandleCommand(msg, commandPath)
    default:
        log.Printf("OzonCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
    }
}
