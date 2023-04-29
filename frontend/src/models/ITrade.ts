import {AdminInterface} from "./IAdmin";
import {UserInterface} from "./IUser";
import {StatusInterface} from "./IStatus";
import {DrugInterface} from "./IDrug";

export interface TradeInterface {
	IDtrade: 		number,
    QUANTITY:   number,
    COST: 		number,
	Total: 	    number,
	Evidence:  	number,
    TradeTime:   Date | null,


    AdminID:    number,
    Admin:      AdminInterface,

    UserID:     number,
    User:       UserInterface,

    DrugID:     number,
    Drug:       DrugInterface,

    StatusID:   number,
    Status:     StatusInterface,
}