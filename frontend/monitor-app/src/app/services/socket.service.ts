import { LogService } from './log.service';
import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { Log } from '../models/log';
import { ProcessesService } from './processes.service';
import { handleAutoChangeDetectionStatus } from '@angular/cdk/testing';

@Injectable({
  providedIn: 'root'
})
export class SocketService {
  private log_socket: WebSocket;
  private process_socket: WebSocket;
  private root_address: string = 'localhost:8080';


  constructor(private logService: LogService, private proccessesService: ProcessesService) {
    this.log_socket = new WebSocket('ws://'+ this.root_address +'/ws');
    this.process_socket = new WebSocket('ws://'+ this.root_address +'/wsp');

    this.log_socket.onmessage = (event) => {
      const log: Log = JSON.parse(event.data);
      logService.addLog(log);
      console.log(log);
    };
    this.process_socket.onmessage = (event) => {
      const conn_message = JSON.parse(event.data);
      proccessesService.handleConnectionMessage(conn_message);
    };
  }
}
