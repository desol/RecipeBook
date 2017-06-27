import { Router } from '@angular/router';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { CanActivate } from '@angular/router';

@Injectable()
export class RootPathGuard implements CanActivate {
  constructor(private urlRouter: Router) {
  }
  canActivate(): boolean {
    // if (true) {
    //   this.urlRouter.navigate(['Account']);
    // } else {
    //   this.urlRouter.navigate(['Login']);
    // }

    return true;
  }
}
