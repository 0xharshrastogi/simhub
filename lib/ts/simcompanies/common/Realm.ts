import { RealmType } from "./RealmType";
import { ResourceBase } from "./ResourceBase";

export abstract class Realm {
  abstract realmType: RealmType;

  abstract fetchResourceList(): Promise<ResourceBase[]>;
}
