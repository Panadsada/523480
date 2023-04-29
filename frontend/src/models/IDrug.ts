import {AdminInterface} from "./IAdmin";
import {UserInterface} from "./IUser";
import {UnitInterface} from "./IUnit";

export interface DrugInterface {
	IDdrug: 		number,
    Code:       string,
    Name: 		string,
	Amount: 	number,
	Price:  	number,

    AdminID:    number,
    Admin:      AdminInterface,

    UserID:     number,
    User:       UserInterface,

    UnitID:     number,
    Unit:       UnitInterface,
}