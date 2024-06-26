import { Injectable } from '@angular/core';
import { Process } from '../models/process';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ProcessesService {

  private processes: Process[] = [

  ];

  private processesSubject = new BehaviorSubject<Process[]>(this.processes);
  processes$: Observable<Process[]> = this.processesSubject.asObservable();

  addNewProcess(process: Process): void {
    this.processes.push(process);
    this.processesSubject.next(this.processes);
  }

  getProccesByName(name: string): Process| undefined{
    return this.processes.find(p => p.name === name);
  }

  handleConnectionMessage(conn_message: Process) {
    const process = this.getProccesByName(conn_message.name);
    if (process) {
      process.monitoring = conn_message.monitoring;
     if (conn_message.connected != process.connected){
      process.connected = conn_message.connected;
      this.handleProcessConnection(process);
     }

    }else{
      this.addNewProcess(conn_message)
    }
    this.processesSubject.next(this.processes);
  }

  deleteProcess(processName: string): void {
    this.processes = this.processes.filter(process => process.name !== processName);
    this.processesSubject.next(this.processes);
  }

  handleProcessConnection(process: Process): void {
      //if connected connect message if dissconected dissconect message
  }

  disconnectProcess(process: Process){
    process.connected = false;
    this.processesSubject.next(this.processes);
  }

  toggleMonitoring(process: Process){
    process.monitoring = !process.monitoring;
    //api logic here send message to the backend to stop sending
  }
}
