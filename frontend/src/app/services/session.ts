import { Observable } from 'rxjs/Observable';
import { Router, NavigationEnd, NavigationStart } from '@angular/router';
import { Injectable, EventEmitter } from '@angular/core';

@Injectable()
export class Session {

  PageLoading: EventEmitter<boolean> = new EventEmitter<boolean>(false);

  constructor(private ngRouter: Router) {
    this.ngRouter.events.filter(event => event instanceof NavigationEnd).subscribe((navEnd) => {
      this.ShowLoading();
    });
    this.ngRouter.events.filter(event => event instanceof NavigationStart).subscribe((navStart) => {
      this.HideLoading();
    });
  }

  private ShowLoading() {
    this.PageLoading.emit(true);
  }

    private HideLoading() {
    this.PageLoading.emit(false);
  }
}
