import { Observable } from 'rxjs/Observable';
import { Router, NavigationEnd, NavigationStart } from '@angular/router';
import { Injectable, EventEmitter } from '@angular/core';

@Injectable()
export class Session {

  LoggedInStatus: EventEmitter<boolean> = new EventEmitter<boolean>(false);
  private LoggedIn = false;

  RequestsLoadingCounter: EventEmitter<number> = new EventEmitter<number>(false);
  private RequestsLoading = 0;

  private AddLoading() {
    this.RequestsLoading += 1;
    this.RequestsLoadingCounter.emit(this.RequestsLoading);
  }

  private HideLoading() {
    this.RequestsLoading <= 1 ? 0 : this.RequestsLoading -= 1;
    this.RequestsLoadingCounter.emit(this.RequestsLoading);
  }

  private UserLoggedIn() {
    this.LoggedIn = true;
    this.LoggedInStatus.emit(this.LoggedIn);
  }

  private UserLoggedOut() {
    this.LoggedIn = false;
    this.LoggedInStatus.emit(this.LoggedIn);
  }
}
