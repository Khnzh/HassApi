package bot

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var currentInlineMarkup string

var setTimeKeyboard = inlineKeyboardMarkup{Type: "setTime", Markup: tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Часы", "hh"),
		tgbotapi.NewInlineKeyboardButtonData("Минуты", "mi"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Установить", "ok"),
	),
)}

var hoursKeyboard = inlineKeyboardMarkup{Type: "hours", Markup: tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "1"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("7", "7"),
		tgbotapi.NewInlineKeyboardButtonData("8", "8"),
		tgbotapi.NewInlineKeyboardButtonData("9", "9"),
		tgbotapi.NewInlineKeyboardButtonData("10", "10"),
		tgbotapi.NewInlineKeyboardButtonData("11", "11"),
		tgbotapi.NewInlineKeyboardButtonData("12", "12"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("13", "13"),
		tgbotapi.NewInlineKeyboardButtonData("14", "14"),
		tgbotapi.NewInlineKeyboardButtonData("15", "15"),
		tgbotapi.NewInlineKeyboardButtonData("16", "16"),
		tgbotapi.NewInlineKeyboardButtonData("17", "17"),
		tgbotapi.NewInlineKeyboardButtonData("18", "18"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("19", "19"),
		tgbotapi.NewInlineKeyboardButtonData("20", "20"),
		tgbotapi.NewInlineKeyboardButtonData("21", "21"),
		tgbotapi.NewInlineKeyboardButtonData("22", "22"),
		tgbotapi.NewInlineKeyboardButtonData("23", "23"),
		tgbotapi.NewInlineKeyboardButtonData("00", "00"),
	),
)}

var minutesKeyboard = inlineKeyboardMarkup{Type: "minutes",
	Markup: tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "2"),
			tgbotapi.NewInlineKeyboardButtonData("3", "3"),
			tgbotapi.NewInlineKeyboardButtonData("4", "4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("6", "6"),
			tgbotapi.NewInlineKeyboardButtonData("7", "7"),
			tgbotapi.NewInlineKeyboardButtonData("8", "8"),
			tgbotapi.NewInlineKeyboardButtonData("9", "9"),
			tgbotapi.NewInlineKeyboardButtonData("10", "10"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("11", "11"),
			tgbotapi.NewInlineKeyboardButtonData("12", "12"),
			tgbotapi.NewInlineKeyboardButtonData("13", "13"),
			tgbotapi.NewInlineKeyboardButtonData("14", "14"),
			tgbotapi.NewInlineKeyboardButtonData("15", "15"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("16", "16"),
			tgbotapi.NewInlineKeyboardButtonData("17", "17"),
			tgbotapi.NewInlineKeyboardButtonData("18", "18"),
			tgbotapi.NewInlineKeyboardButtonData("19", "19"),
			tgbotapi.NewInlineKeyboardButtonData("20", "20"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("21", "21"),
			tgbotapi.NewInlineKeyboardButtonData("22", "22"),
			tgbotapi.NewInlineKeyboardButtonData("23", "23"),
			tgbotapi.NewInlineKeyboardButtonData("24", "24"),
			tgbotapi.NewInlineKeyboardButtonData("25", "25"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("26", "26"),
			tgbotapi.NewInlineKeyboardButtonData("27", "27"),
			tgbotapi.NewInlineKeyboardButtonData("28", "28"),
			tgbotapi.NewInlineKeyboardButtonData("29", "29"),
			tgbotapi.NewInlineKeyboardButtonData("30", "30"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("31", "31"),
			tgbotapi.NewInlineKeyboardButtonData("32", "32"),
			tgbotapi.NewInlineKeyboardButtonData("33", "33"),
			tgbotapi.NewInlineKeyboardButtonData("34", "34"),
			tgbotapi.NewInlineKeyboardButtonData("35", "35"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("36", "36"),
			tgbotapi.NewInlineKeyboardButtonData("37", "37"),
			tgbotapi.NewInlineKeyboardButtonData("38", "38"),
			tgbotapi.NewInlineKeyboardButtonData("39", "39"),
			tgbotapi.NewInlineKeyboardButtonData("40", "40"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("41", "41"),
			tgbotapi.NewInlineKeyboardButtonData("42", "42"),
			tgbotapi.NewInlineKeyboardButtonData("43", "43"),
			tgbotapi.NewInlineKeyboardButtonData("44", "44"),
			tgbotapi.NewInlineKeyboardButtonData("45", "45"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("46", "46"),
			tgbotapi.NewInlineKeyboardButtonData("47", "47"),
			tgbotapi.NewInlineKeyboardButtonData("48", "48"),
			tgbotapi.NewInlineKeyboardButtonData("49", "49"),
			tgbotapi.NewInlineKeyboardButtonData("50", "50"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("51", "51"),
			tgbotapi.NewInlineKeyboardButtonData("52", "52"),
			tgbotapi.NewInlineKeyboardButtonData("53", "53"),
			tgbotapi.NewInlineKeyboardButtonData("54", "54"),
			tgbotapi.NewInlineKeyboardButtonData("55", "55"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("56", "56"),
			tgbotapi.NewInlineKeyboardButtonData("57", "57"),
			tgbotapi.NewInlineKeyboardButtonData("58", "58"),
			tgbotapi.NewInlineKeyboardButtonData("59", "59"),
			tgbotapi.NewInlineKeyboardButtonData("00", "00"),
		)),
}

type inlineKeyboardMarkup struct {
	Type   string
	Markup tgbotapi.InlineKeyboardMarkup
}

func (i inlineKeyboardMarkup) getType() string {
	return i.Type
}

type inlineKeyboard interface {
	getType() string
}

func getType(i inlineKeyboard) string {
	return i.getType()
}

type inlineMessage struct {
	Blocks  bool
	Message *tgbotapi.Message
}

func (i inlineMessage) init(m tgbotapi.Message) *inlineMessage {
	i.Message = &m
	i.Blocks = true
	return &i
}

func (i inlineMessage) getId() int {
	return i.Message.MessageID
}

func (i inlineMessage) getChatId() int64 {
	return i.Message.Chat.ID
}

func InlKeyUpdateParams(i *inlineMessage, m inlineKeyboardMarkup) tgbotapi.Params {
	params := tgbotapi.Params{}
	params.AddNonZero64("chat_id", i.getChatId())
	params.AddNonZero("message_id", i.getId())
	params.AddInterface("reply_markup", m.Markup)
	return params
}

// runSudo runs a command via sudo, passing the password via stdin.
func runSudo(password string, args ...string) ([]byte, error) {
	cmd := exec.Command("sudo", append([]string{"-S"}, args...)...)
	cmd.Stdin = strings.NewReader(password + "\n")
	return cmd.CombinedOutput()
}

// applySchedule writes the shutdown script, makes it executable, and updates the crontab.
func applySchedule(sudoPass, hh, mm string) error {
	var day string
	if time.Now().Hour() >= 6 {
		day = "tomorrow "
	}

	// Write /usr/local/bin/daily-shutdown.sh
	scriptContent := fmt.Sprintf("sudo rtcwake -m no -t $(date -d \"%v%v:%v\" +%%s)\n", day, hh, mm)
	writeCmd := exec.Command("sudo", "-S", "tee", "/usr/local/bin/daily-shutdown.sh")
	writeCmd.Stdin = strings.NewReader(sudoPass + "\n" + scriptContent)
	if out, err := writeCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("write script: %w (output: %s)", err, out)
	}

	if out, err := runSudo(sudoPass, "chmod", "+x", "/usr/local/bin/daily-shutdown.sh"); err != nil {
		return fmt.Errorf("chmod script: %w (output: %s)", err, out)
	}

	// Read current root crontab
	listCmd := exec.Command("sudo", "-S", "crontab", "-l")
	listCmd.Stdin = strings.NewReader(sudoPass + "\n")
	crontabOut, _ := listCmd.Output()

	// Remove the last line (previous schedule entry) and append the new one
	lines := []string{}
	if len(crontabOut) > 0 {
		existing := strings.Split(strings.TrimRight(string(crontabOut), "\n"), "\n")
		if len(existing) > 1 {
			lines = existing[:len(existing)-1]
		}
	}
	lines = append(lines, fmt.Sprintf("%v %v * * * /usr/local/bin/daily-shutdown.sh", mm, hh))
	newCrontab := strings.Join(lines, "\n") + "\n"

	setCmd := exec.Command("sudo", "-S", "crontab", "-")
	setCmd.Stdin = strings.NewReader(sudoPass + "\n" + newCrontab)
	if out, err := setCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("set crontab: %w (output: %s)", err, out)
	}

	return nil
}

func ListenAndServe() {
	// Load .env if present; fall back to environment variables if not.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	var InlineKeyboard *inlineMessage
	var hh string
	var mm string

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = setTimeKeyboard.Markup
				currentInlineMarkup = getType(setTimeKeyboard)
				hh = ""
				mm = ""
			}

			if m, err := bot.Send(msg); err != nil {
				log.Println("send:", err)
			} else {
				InlineKeyboard = inlineMessage{}.init(m)
			}

		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				log.Println("callback:", err)
				continue
			}

			switch update.CallbackQuery.Data {
			case "hh":
				if currentInlineMarkup != getType(hoursKeyboard) {
					params := InlKeyUpdateParams(InlineKeyboard, hoursKeyboard)
					if _, err := bot.MakeRequest("editMessageReplyMarkup", params); err != nil {
						log.Println("edit markup:", err)
					} else {
						currentInlineMarkup = getType(hoursKeyboard)
					}
				}

			case "mi":
				if currentInlineMarkup != getType(minutesKeyboard) {
					params := InlKeyUpdateParams(InlineKeyboard, minutesKeyboard)
					if _, err := bot.MakeRequest("editMessageReplyMarkup", params); err != nil {
						log.Println("edit markup:", err)
					} else {
						currentInlineMarkup = getType(minutesKeyboard)
					}
				}

			case "ok":
				var messageText string
				if hh == "" || mm == "" {
					messageText = "Choose both hours and minutes values to set lights toggle timing!"
				} else {
					sudoPass := os.Getenv("SUDO_PASSWORD")
					hassToken := os.Getenv("HASS_TOKEN")
					errs := []string{}

					if err := applySchedule(sudoPass, hh, mm); err != nil {
						log.Println("applySchedule:", err)
						errs = append(errs, err.Error())
					}

					if err := updateHassAutomationTime(hassToken, hh, mm); err != nil {
						log.Println("updateHass:", err)
						errs = append(errs, err.Error())
					}

					if len(errs) > 0 {
						messageText = fmt.Sprintf("Scheduled %v:%v, but errors occurred: %v", hh, mm, strings.Join(errs, "; "))
					} else {
						messageText = fmt.Sprintf("Lights will toggle at %v:%v", hh, mm)
					}
				}

				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, messageText)
				if _, err := bot.Send(msg); err != nil {
					log.Println("send:", err)
				}

			default:
				if currentInlineMarkup != getType(setTimeKeyboard) {
					if currentInlineMarkup == getType(hoursKeyboard) {
						if len(update.CallbackQuery.Data) < 2 {
							hh = "0" + update.CallbackQuery.Data
						} else {
							hh = update.CallbackQuery.Data
						}
					}
					if currentInlineMarkup == getType(minutesKeyboard) {
						if len(update.CallbackQuery.Data) < 2 {
							mm = "0" + update.CallbackQuery.Data
						} else {
							mm = update.CallbackQuery.Data
						}
					}
					params := InlKeyUpdateParams(InlineKeyboard, setTimeKeyboard)
					if _, err := bot.MakeRequest("editMessageReplyMarkup", params); err != nil {
						log.Println("edit markup:", err)
					} else {
						currentInlineMarkup = getType(setTimeKeyboard)
					}
				}
			}
		}
	}
}
