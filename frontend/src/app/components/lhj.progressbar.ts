import {Component} from '@angular/core';
import { Session } from '../services/session'

@Component({
  selector: 'app-lhj-progressbar',
  templateUrl: '../pages/lhj.progressbar.html',
  styleUrls: ['../styles/lhj.progressbar.css'],
})
export class ProgressBarComponent {
  loading = false;
  constructor(session: Session) {
    session.PageLoading.subscribe(
      (load) => {
        this.loading = load;
      }
    )
  }
}
