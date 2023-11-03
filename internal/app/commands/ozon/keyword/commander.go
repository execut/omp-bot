package keyword

import (
    "github.com/ozonmp/omp-bot/internal/service/ozon/keyword"
    "log"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/ozonmp/omp-bot/internal/app/path"
)

type OzonKeywordCommander struct {
    bot            *tgbotapi.BotAPI
    keywordService *keyword.Service
}

func NewOzonKeywordCommander(
    bot *tgbotapi.BotAPI,
) *OzonKeywordCommander {
    subdomainService := keyword.NewService()

    return &OzonKeywordCommander{
        bot:            bot,
        keywordService: subdomainService,
    }
}

func (c *OzonKeywordCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
    switch callbackPath.CallbackName {
    case "list":
        c.CallbackList(callback, callbackPath)
    default:
        log.Printf("OzonKeywordCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
    }
}

func (c *OzonKeywordCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
    switch commandPath.CommandName {
    case "help":
        c.Help(msg)
    case "list":
        c.List(msg)
    case "get":
        c.Get(msg)
    case "delete":
        c.Delete(msg)
    default:
        c.Default(msg)
    }
}
