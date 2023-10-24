package proto

import "encoding/json"

// Request -- запрос клиента к серверу.
type Request struct {
	// Поле Command может принимать три значения:
	// * "quit" - прощание с сервером (после этого сервер рвёт соединение);
	// * "calc" - передача нового значения на сервер;
	Command string `json:"command"`

	Data *json.RawMessage `json:"data"`
}

// Response -- ответ сервера клиенту.
type Response struct {
	// Поле Status может принимать три значения:
	// * "ok" - успешное выполнение команды "quit" или "add";
	// * "failed" - в процессе выполнения команды произошла ошибка;
	// * "result" - ответ вычислен
	Status string `json:"status"`

	// Если Status == "failed", то в поле Data находится сообщение об ошибке.
	// Если Status == "result", в поле Data должен лежать Answer
	// В противном случае, поле Data пустое.
	Data *json.RawMessage `json:"data"`
}

// Conditions -- условия броска
type Conditions struct {
	// Angle -- угол броска
	Angle string `json:"ang"`

	// Speed -- скорость броска
	Speed string `json:"spd"`
}

// Answer -- ответ на задачу.
type Answer struct {
	// Численное представление ответа
	Result string `json:"res"`
}
