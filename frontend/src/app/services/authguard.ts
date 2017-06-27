import { Router } from '@angular/router';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { CanActivate } from '@angular/router';

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(private urlRouter: Router) {
  }
  canActivate(): boolean {
    let ReturnVal = true;
    // if (Cookie.get('token') && Cookie.get('username')) {
    //   ReturnVal = true;
    //   this.sessionManager.UserLoggedInOut(true);
    // } else {
    //   this.Toast.error('Your Authorization Failed.', 'Redirected To Login');
    //   this.sessionManager.UserLoggedInOut(false);
    //   this.urlRouter.navigate(['Login'], { relativeTo: null });
    // }
    return ReturnVal;
  }
}
