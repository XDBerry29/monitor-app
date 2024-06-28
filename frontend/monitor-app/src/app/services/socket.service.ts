import { LogService } from './log.service';
import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { Log } from '../models/log';
import { ProcessesService } from './processes.service';
import { handleAutoChangeDetectionStatus } from '@angular/cdk/testing';
import { v4 as uuidv4 } from 'uuid';
import { LogFilterMsg } from '../models/log-filter-msg';
import { ProcessFilterMsg } from '../models/process-filter-msg';

@Injectable({
  providedIn: 'root'
})
export class SocketService {
  private log_socket: WebSocket;
  private process_socket: WebSocket;
  private root_address: string = 'localhost:8080';


  constructor(private logService: LogService, private proccessesService: ProcessesService) {
    const clientId = uuidv4();

    this.log_socket = new WebSocket('ws://'+ this.root_address +'/ws');
    this.log_socket.onopen = () => {
      this.log_socket.send(JSON.stringify({ type: 'init', clientId }));
    };

    this.process_socket = new WebSocket('ws://'+ this.root_address +'/wsp');
    this.process_socket.onopen = () => {
      this.process_socket.send(JSON.stringify({ type: 'init', clientId }));
    };

    this.log_socket.onmessage = (event) => {
      const log: Log = JSON.parse(event.data);
      logService.addLog(log);
      //console.log(log);
    };
    this.process_socket.onmessage = (event) => {
      const conn_message = JSON.parse(event.data);
      proccessesService.handleConnectionMessage(conn_message);
    };
  }

  sendLogFilterMessage(filter: LogFilterMsg): void {
    if (this.log_socket.readyState === WebSocket.OPEN) {
      this.log_socket.send(JSON.stringify(filter));
    } else {
      console.error('Log WebSocket connection is not open');
    }
  }

  sendProcessFilterMessage(filter: ProcessFilterMsg): void {
    if (this.process_socket.readyState === WebSocket.OPEN) {
      this.process_socket.send(JSON.stringify(filter));
    } else {
      console.error('Process WebSocket connection is not open');
    }
  }



}
