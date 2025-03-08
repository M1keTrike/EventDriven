import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { WebSocketService } from '../../services/websocket.service';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.css']
})
export class ChatComponent implements OnInit, OnDestroy {
  messages: any[] = [];
  newMessage: string = '';
  private wsSubscription!: Subscription;

  constructor(private wsService: WebSocketService) {}

  ngOnInit() {
    this.wsSubscription = this.wsService.getMessages().subscribe(message => {
      this.messages.push(message);
    });
  }

  sendMessage() {
    if (this.newMessage.trim() !== '') {
      this.wsService.sendMessage({ sender: 'User1', content: this.newMessage });
      this.newMessage = '';
    }
  }

  ngOnDestroy() {
    this.wsSubscription.unsubscribe();
    this.wsService.close();
  }
}
