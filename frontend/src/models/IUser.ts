import {StatusInterface} from "./IStatus";
import {GenderInterface} from "./IGender";
import {TitleInterface} from "./ITitle";

export interface UserInterface {
	ID:         number,
	Name: 		string,
    Personal:   string,
	Email:      string,
	Password: 	string,
    Tel:        string,
    BirthdayTime:   Date | null,
	Role: 		string,

    StatusID:   number,
    Status:     StatusInterface,

    GenderID:   number,
    Gender:     GenderInterface,

    TitleID:    number,
    Title:      TitleInterface,
}