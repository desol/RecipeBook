import { Http, Response } from '@angular/http';

export class Client {
  protected readonly BaseURL = 'http://localhost:4200/api';

  constructor(protected http: Http) {}
}
