import { Injectable } from '@angular/core';
import { WebSocketSubject } from 'rxjs/webSocket';
import { Observable } from 'rxjs';

const WEBSOCKET_URL = 'ws://localhost:8080/ws';

@Injectable({
  providedIn: 'root',
})
export class WebSocketService {
  private socket$: WebSocketSubject<any>;

  constructor() {
    this.socket$ = new WebSocketSubject(WEBSOCKET_URL);
  }


  sendMessage(message: any): void {
    this.socket$.next(message);
  }


  getMessages(): Observable<any> {
    return this.socket$.asObservable();
  }


  close(): void {
    this.socket$.complete();
  }
}
