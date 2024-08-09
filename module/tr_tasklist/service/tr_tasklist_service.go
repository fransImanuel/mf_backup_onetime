package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mf_backup_onetime/constant"
	"mf_backup_onetime/dto"
	"mf_backup_onetime/module/ms_destination"
	"mf_backup_onetime/module/ms_questions_ac"
	"mf_backup_onetime/module/ms_user"
	"mf_backup_onetime/module/tr_tasklist"
	"mf_backup_onetime/module/tr_tasklist/model"
	"mf_backup_onetime/pkg"
	"mf_backup_onetime/schemas"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/signintech/gopdf"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TRTasklistService struct {
	TRTasklist           tr_tasklist.Repository
	MSQuestionRepository ms_questions_ac.Repository
	MSDestination        ms_destination.Repository
	MSUsers              ms_user.Repository
}

func InitTRTasklistService(tr_tasklistRepository tr_tasklist.Repository,
	MSQuestionRepository ms_questions_ac.Repository, MSDestination ms_destination.Repository, MSUsers ms_user.Repository) tr_tasklist.Service {
	return &TRTasklistService{
		TRTasklist:           tr_tasklistRepository,
		MSQuestionRepository: MSQuestionRepository,
		MSDestination:        MSDestination,
		MSUsers:              MSUsers,
	}
}

func (s *TRTasklistService) BulkExportOneTimeService() []string {
	return s.TRTasklist.BulkMongoExportOneTime()
}

func (s *TRTasklistService) ExportPDFTasklist(ctx context.Context, input dto.GetTasklistByID) (*string, string, schemas.SchemaDatabaseError) {

	TasklistId, err := primitive.ObjectIDFromHex(input.TasklistId)
	if err != nil {
		log.Errorln("Error TasklistId ==> ", err)
	}

	TRTaskList, errFind := s.TRTasklist.GetTasklistByIdRepository(TasklistId)
	if errFind.Error != nil {
		return nil, "", errFind
	}
	fmt.Printf("TRTaskList %+v\n ", TRTaskList)

	Destination, err := s.MSDestination.GetById(TRTaskList.DestinationId)
	if err != nil {
		log.Errorln("Err GetDestination ", err)
		return nil, "", schemas.SchemaDatabaseError{Error: err, Code: 400}
	}

	GetMSQuestionSurvey, err := s.MSQuestionRepository.GetMSQuestionSurveys(ctx)
	if err != nil {
		log.Errorln("Err GetMSQuestionSurvey ", err)
		return nil, "", schemas.SchemaDatabaseError{Error: err, Code: 400}
	}
	Users, errFind := s.MSUsers.UserById(TRTaskList.AssignedUserId)
	if errFind.Error != nil {
		log.Errorln("Err GetMSQuestionSurvey ", errFind.Error)
		return nil, "", errFind
	}

	DataPDF := dto.ReportPDFList{}
	var ListQuestionAnswer []*dto.QuestionAnswer

	if len(TRTaskList.TasklistDetail) > 0 && TRTaskList.TasklistDetail[len(TRTaskList.TasklistDetail)-1] != nil {

		for _, survey := range TRTaskList.TasklistDetail[0].ResultSurvey {
			var QuestionAnswer dto.QuestionAnswer
			var Question dto.Question
			var AnswerQuestion model.ResultSurvey
			for _, question := range GetMSQuestionSurvey.Question {
				if question.Id == survey.QuestionId {
					Question = question
				}
			}
			AnswerQuestion = survey
			QuestionAnswer.Question = Question
			QuestionAnswer.Answer = AnswerQuestion
			ListQuestionAnswer = append(ListQuestionAnswer, &QuestionAnswer)
		}
		if TRTaskList.DestinationId == Destination.ID && Users.ID == TRTaskList.AssignedUserId {

			DataPDF.DestinationName = Destination.Name
			DataPDF.DestinationCode = Destination.Code
			DataPDF.DestinationAddress = Destination.Address

			DataPDF.AssignedUser = Users.FullName

			DataPDF.TasklistId = TRTaskList.Id
			DataPDF.Duration = TRTaskList.TasklistDetail[len(TRTaskList.TasklistDetail)-1].DurationTime
			DataPDF.SurveyTime = TRTaskList.TasklistDetail[len(TRTaskList.TasklistDetail)-1].SurveyTime
			DataPDF.ScheduleVisit = TRTaskList.ScheduleVisit
			DataPDF.QuestionAnswer = ListQuestionAnswer

			switch TRTaskList.StatusId {
			case 0:
				DataPDF.Status = "OPEN"
				break
			case 1:
				DataPDF.Status = "WAITING SYNC"
				break
			case 2:
				DataPDF.Status = "WAITING APPROVAL"
				break
			case 3:
				DataPDF.Status = "OPEN SCHEDULE"
				break
			case 4:
				DataPDF.Status = "DONE"
				break
			case 5:
				DataPDF.Status = "DELAYED"
				break
			case 6:
				DataPDF.Status = "UNDONE"
				break
			case 7:
				DataPDF.Status = "MANUAL"

			}

		}

	}

	base64, fname, err := s.GenerateFilePDF(ctx, DataPDF)
	if err != nil {
		log.Errorln("Err Generate PDF")
		return nil, "", schemas.SchemaDatabaseError{
			Error: err,
			Code:  500,
		}
	}

	return base64, fname, schemas.SchemaDatabaseError{}
}

func (s *TRTasklistService) GenerateFilePDF(ctx context.Context, Data dto.ReportPDFList) (*string, string, error) {
	var MaxHeight, HeightImage, WidthImage, FontSize float64
	MaxHeight = 842
	HeightImage = 310
	WidthImage = 225
	FontSize = 12

	var CurrentHeight float64

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	//pdf.ImageByHolder()

	//err := pdf.AddTTFFont("loma", "LiberationSerif-Regular.ttf")
	// err := pdf.AddTTFFontWithOption("loma", "assets/font/LiberationSerif-Regular.ttf", gopdf.TtfOption{Style: 0})
	// if err != nil {
	// 	log.Println(err.Error())
	// 	// return nil, "", err
	// }
	// err = pdf.AddTTFFontWithOption("loma", "assets/font/LiberationSerif-Bold.ttf", gopdf.TtfOption{Style: 2})
	// if err != nil {
	// 	log.Println(err.Error())
	// 	// return nil, "", err
	// }

	// err = pdf.SetFont("loma", "", FontSize)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	// return nil, "", err
	// }
	pdf.SetFontSize(FontSize)
	//var MaxWidth = 595

	title := &gopdf.Rect{
		W: 100,
		H: 20}
	value := &gopdf.Rect{
		W: 100,
		H: 20}

	pdf.SetXY(30, 20)
	pdf.Cell(title, "ATM ID")
	pdf.Cell(value, fmt.Sprintf(": %v", Data.DestinationCode))
	pdf.SetXY(30, 40)
	pdf.Cell(title, "ATM Name")
	pdf.Cell(value, fmt.Sprintf(": %v", Data.DestinationName))
	pdf.SetXY(30, 60)
	pdf.Cell(title, "Address :")
	pdf.SetXY(30, 80)
	pdf.Cell(value, fmt.Sprintf("%v", Data.DestinationAddress))
	pdf.SetXY(30, 100)
	pdf.Cell(title, "Submitted By")
	pdf.Cell(value, fmt.Sprintf(": %v", Data.AssignedUser))
	pdf.SetXY(30, 120)
	pdf.Cell(title, "Date")
	loc, _ := time.LoadLocation("Asia/Jakarta")

	pdf.Cell(value, fmt.Sprintf(": %v", Data.SurveyTime.In(loc).Format("02/01/2006 15:04:05")))
	pdf.SetXY(30, 140)
	pdf.Cell(title, "Status")
	pdf.Cell(value, fmt.Sprintf(": %v", Data.Status))
	pdf.SetXY(30, 160)
	pdf.Cell(title, "Duration")
	pdf.Cell(value, fmt.Sprintf(": %v", Data.Duration))
	pdf.Line(10, 180, 585, 180)
	CurrentHeight = 200

	for i, Question := range Data.QuestionAnswer {
		if CurrentHeight >= (MaxHeight - 60) {
			pdf.AddPage()
			CurrentHeight = 30
		}
		pdf.SetXY(30, CurrentHeight)

		var pertanyaan = fmt.Sprintf("%v. %v", i+1, Question.Question.Label)
		words := strings.Fields(pertanyaan)
		//fmt.Println(words, len(words))
		var text, text2, text3 string

		for _, word := range words {
			if (len(text) + 1) <= 55 {
				text += fmt.Sprintf("%v ", word)
			} else if (len(text2) + 1) <= 55 {
				text2 += fmt.Sprintf("%v ", word)
			} else {
				text3 += fmt.Sprintf("%v ", word)
			}
		}

		err = pdf.SetFont("loma", "B", FontSize)
		if err != nil {
			log.Print(err.Error())
			return nil, "", err
		}
		pdf.Cell(nil, text)
		CurrentHeight += 20
		if len(text2) > 1 {
			pdf.SetXY(40, CurrentHeight)
			pdf.Cell(nil, text2)
			CurrentHeight += 20
		}
		if len(text3) > 1 {
			pdf.SetXY(40, CurrentHeight)
			pdf.Cell(nil, text3)
			CurrentHeight += 20
		}
		err = pdf.SetFont("loma", "", FontSize)
		if err != nil {
			log.Print(err.Error())
			return nil, "", err
		}
		if Question.Question.TypeField == "radio-button" {
			for _, item := range Question.Question.Item {
				if CurrentHeight >= (MaxHeight - 60) {
					pdf.AddPage()
					CurrentHeight = 30
				}
				pdf.SetXY(40, CurrentHeight)
				intVar, _ := strconv.Atoi(fmt.Sprintf("%v", Question.Answer.Answer.ResultItem[0].Value))
				if intVar == item.Value {
					pdf.Cell(&gopdf.Rect{W: 15, H: 20}, "")
					pdf.Image("assets/img/radio-on-button.png", 40, CurrentHeight, &gopdf.Rect{
						W: 15,
						H: 15})
					pdf.Cell(&gopdf.Rect{W: 15, H: 20}, fmt.Sprintf(" %v", item.Key))
				} else {
					pdf.Cell(&gopdf.Rect{W: 15, H: 20}, "")
					pdf.Image("assets/img/radio-button.png", 40, CurrentHeight, &gopdf.Rect{
						W: 15,
						H: 15})
					pdf.Cell(&gopdf.Rect{W: 15, H: 20}, fmt.Sprintf(" %v", item.Key))
				}
				CurrentHeight += 20

			}
		} else if Question.Question.TypeField == "textbox" {
			if CurrentHeight >= (MaxHeight - 60) {
				pdf.AddPage()
				CurrentHeight = 30
			}
			pdf.SetXY(40, CurrentHeight)
			pdf.Cell(&gopdf.Rect{W: 100, H: 20}, fmt.Sprintf(" %v", *Question.Answer.Answer.Value))
			CurrentHeight += 20
		} else if Question.Question.TypeField == "camera" {
			if CurrentHeight >= (MaxHeight - (60 + 200)) {
				pdf.AddPage()
				CurrentHeight = 30
			}
			fileName := fmt.Sprintf("%s ", Question.Answer.Id.Hex())

			URL := ""
			if len(Question.Answer.Answer.ResultItem) > 0 {
				URL = fmt.Sprintf("http://10.254.212.5:81/%v", Question.Answer.Answer.ResultItem[0].Value)
			}
			// URL := fmt.Sprintf("%v/%v", pkg.GodotEnv("DOMAIN_IMAGE"), Question.Answer.Answer.ResultItem[0].Value)
			fmt.Println("===================================================================")
			fmt.Println("URL DOWNLOAD:", URL)
			fmt.Println("===================================================================")
			err = DownloadFile(URL, fileName)
			if err == nil {
				errImage := pdf.Image(fileName, 40, CurrentHeight, &gopdf.Rect{
					W: WidthImage,
					H: HeightImage})
				CurrentHeight += HeightImage
				CurrentHeight += 20
				if errImage != nil {
					fmt.Println("Error Download ", errImage.Error())
				}

				//err = os.Remove(fileName)
				//if err != nil {
				//	log.Errorln("Error Remove ", err.Error())
				//	return nil, err
				//}
			} else {
				fmt.Println("Error Download ", err.Error())
			}
			//fmt.Printf("File %s downlaod in current working directory", fileName)

		}

		if Question.Answer.Answer.ResultProperty.TakePhoto != nil {
			if CurrentHeight >= (MaxHeight - (60 + 200)) {
				pdf.AddPage()
				CurrentHeight = 30
			}
			fileName := fmt.Sprintf("%s ", Question.Answer.Id.Hex())

			URL := fmt.Sprintf("%v/%v", pkg.GodotEnv("DOMAIN_IMAGE"), *Question.Answer.Answer.ResultProperty.TakePhoto)
			fmt.Println("===================================================================")
			fmt.Println("URL DOWNLOAD:", URL)
			fmt.Println("===================================================================")
			err = DownloadFile(URL, fileName)
			if err == nil {
				errImage := pdf.Image(fileName, 40, CurrentHeight, &gopdf.Rect{
					W: WidthImage,
					H: HeightImage})
				CurrentHeight += HeightImage
				CurrentHeight += 20
				if errImage != nil {
					fmt.Println("Error Download ", errImage.Error())
				}
				//err = os.Remove(fileName)
				//if err != nil {
				//	log.Errorln("Error Remove ", err.Error())
				//	return nil, err
				//}
			} else {
				fmt.Println("Error Download ", err.Error())
			}

		}

	}
	CurrentHeight += 50
	pdf.SetXY(30, CurrentHeight)
	pdf.Cell(title, "TaskListID")
	pdf.Cell(value, fmt.Sprintf(": %v", Data.TasklistId.Hex()))
	CurrentHeight += 50
	pdf.SetXY(30, CurrentHeight)
	pdf.Cell(title, "Schedule Visit")
	pdf.Cell(value, fmt.Sprintf(": %v", Data.ScheduleVisit.Format("02 Jan 2006 15:04:05")))

	var buf bytes.Buffer
	err = pdf.Write(&buf)
	if err != nil {
		log.Errorln(err)
	}

	//!!!!

	//!!!!

	// Mengonversi buffer bytes menjadi string Base64
	base64String := base64.StdEncoding.EncodeToString(buf.Bytes())

	TimeNow := time.Now()
	//loc, _ := time.LoadLocation("Asia/Jakarta")
	year, month, day := Data.ScheduleVisit.In(loc).Date()
	Data.ScheduleVisit.Date()
	location := fmt.Sprintf("%v/%v-%v/%v", constant.STORAGE_PATH_PDF, year, int(month), day)
	if err := os.MkdirAll(location, os.ModePerm); err != nil {
		log.Errorln("MkdirAll ", err)
		return nil, "", err
	}
	filename := fmt.Sprintf("%v/%v_%v_%v-%v-%v-%v.pdf", location, Data.DestinationCode,
		year, int(month), day, Data.Status, TimeNow.UnixMilli())
	pdf.WritePdf(filename)

	f, err := os.OpenFile(fmt.Sprintf("%v-%v.log", year, int(month)),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("ScheduleVisit, %v, ATMID, %v,Status, %v, ATMNAME, %v, Date, %v Location, %v \n", Data.ScheduleVisit.In(loc).Format("2006-01-02 15:04:05"), Data.DestinationCode, Data.Status, Data.DestinationName, TimeNow.Format("2006-01-02 15:04:05"), filename)); err != nil {
		log.Println(err)
	}

	return &base64String, filename, nil
}

func DownloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
