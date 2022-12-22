import { PrefixsInterface } from "./IPrefix";
import { GendersInterface } from "./IGender";
import { RolesInterface } from "./IRole";

export interface OfficerInterface {
    ID: number,
    Name: string,
    Age: number,
    Phone: string,
    Email: string,
    Password: string,
    
    Prefix_ID: number,
    Prefix: PrefixsInterface,

    Gender_ID: number,
    Gender: GendersInterface,

    Role_ID: number,
    Role: RolesInterface,   
  }