package telegram

import (
	"errors"
	"read-adviser-bot/clients/telegram"
	"read-adviser-bot/events"
	"read-adviser-bot/lib/e"
	"read-adviser-bot/storage"
)

type Proc struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

type Meta struct {
	ChatID   int
	Username string
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

func New(client *telegram.Client, storage storage.Storage) *Proc {
	return &Proc{
		tg:      client,
		storage: storage,
	}
}

func (p *Proc) Fetch(limit int) ([]events.Event, error) {
	update, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("Fetch", "can not get events", err)
	}

	if len(update) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(update))

	for _, u := range update {
		res = append(res, event(u))
	}

	p.offset = update[len(update)-1].ID + 1

	return res, nil
}

func (p *Proc) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	default:
		return e.Wrap("processMessage", "can not process message", ErrUnknownEventType)
	}
}

func (p *Proc) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("processMessage", "can not process message", err)
	}

	if err := p.doCmd(event.Text, meta.ChatID, meta.Username); err != nil {
		return e.Wrap("processMessage", "can not process message", err)
	}

	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("meta", "can not get meta", ErrUnknownMetaType)
	}

	return res, nil
}

func event(u telegram.Update) events.Event {
	uType := fetchType(u)

	res := events.Event{
		Type: uType,
		Text: fetchText(u),
	}

	if uType == events.Message {
		res.Meta = Meta{
			ChatID:   u.Message.Chat.ID,
			Username: u.Message.From.Username,
		}
	}
	return res
}

func fetchType(u telegram.Update) events.Type {
	if u.Message == nil {
		return events.Unknown
	}
	return events.Message
}

func fetchText(u telegram.Update) string {
	if u.Message == nil {
		return ""
	}
	return u.Message.Text
}
