import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { User } from '../definitions/auth'

@Component({
  selector: 'app-lhj-login',
  templateUrl: '../pages/lhj.login.html',
  styleUrls: ['../styles/lhj.login.css'],
})
export class LoginComponent {
  username = '';
  password = '';
  user: User;

  constructor(private client: HttpClient) { }

  Login(): void {
    if (this.username.length > 1 && this.password.length > 1) {
      this.client.post<User>('api/auth', JSON.stringify({ email: this.username, password: this.password })).subscribe(user => {
        this.user = user;
        console.log(this.user);
      },
        err => {
          console.log(err);
        });
    }
  }
}
