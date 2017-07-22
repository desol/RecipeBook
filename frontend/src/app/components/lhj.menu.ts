import {Component} from '@angular/core';
import { Session } from '../services/session'

@Component({
  selector: 'app-lhj-menu',
  templateUrl: '../pages/lhj.menu.html',
  styleUrls: ['../styles/lhj.menu.css'],
})
export class MenuComponent {
  loggedIn = false;
  constructor(session: Session) {
    session.LoggedInStatus.subscribe(
      (logInStatus) => {
        this.loggedIn = logInStatus;
      }
    )
  }
}
