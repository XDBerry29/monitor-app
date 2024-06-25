import { LogService } from './log.service';
import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { Log } from '../models/log';

@Injectable({
  providedIn: 'root'
})
export class SocketService {
  private socket: WebSocket;


  constructor(private logService: LogService) {
    this.socket = new WebSocket('ws://localhost:8080/ws');
    this.socket.onmessage = (event) => {
      const log: Log = JSON.parse(event.data);
      logService.addLog(log);
      //console.log(log);
    };
  }
}
