package model

import "time"

//задача
type Task struct {
	Id          uint   `json:"Id"`
	Name        string `json:"Name"`
	Priority    uint   `json:"Priority"`
	Description string `json:"Description"`
	//дата создания таска
	DateCreate time.Time `json:"DateCreate"`
	//дата начала выполнения
	DateStart time.Time `json:"DateStart"`
	//дата предполагаемого окончания выполнения
	DateEnd time.Time `json:"DateEnd"`
	//дата фактического закрытия дела
	DateClose time.Time `json:"DateClose"`
}

//цель
type Target struct {
	Id          uint   `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	//дата создания цели
	DateCreate time.Time `json:"DateCreate"`
	//дата дедлайна
	DeadLine time.Time `json:"DeadLine"`
	//дата фактического достижения цели
	DateClose time.Time `json:"DateClose"`
	//связь с задачами 1-n
	Tasks []Task `json:"Tasks"`
}
