import { Injectable } from '@angular/core';
import { Process } from '../models/process';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ProcessesService {

  private processes: Process[] = [
    {
      name: 'API', connected: false, monitoring: false,
      time: ''
    },
    {
      name: 'Web', connected: true, monitoring: true,
      time: ''
    },
    {
      name: 'Worker', connected: true, monitoring: true,
      time: ''
    },
    {
      name: 'Scheduler', connected: true, monitoring: true,
      time: ''
    },
    {
      name: 'Daemon', connected: true, monitoring: true,
      time: ''
    },
    {
      name: 'Cron', connected: true, monitoring: true,
      time: ''
    },
    {
      name: 'All', connected: true, monitoring: true,
      time: ''
    },{
      name: 'All', connected: true, monitoring: true,
      time: ''
    },{
      name: 'All', connected: true, monitoring: true,
      time: ''
    },{
      name: 'All', connected: true, monitoring: true,
      time: ''
    },{
      name: 'All', connected: true, monitoring: true,
      time: ''
    },{
      name: 'All', connected: true, monitoring: true,
      time: ''
    },{
      name: 'All', connected: true, monitoring: true,
      time: ''
    },{
      name: 'All', connected: true, monitoring: true,
      time: ''
    },{
      name: 'All', connected: true, monitoring: true,
      time: ''
    },{
      name: 'All', connected: true, monitoring: true,
      time: ''
    },{
      name: 'All', connected: true, monitoring: true,
      time: ''
    },{
      name: 'All', connected: true, monitoring: true,
      time: ''
    },
  ];

  private processesSubject = new BehaviorSubject<Process[]>(this.processes);
  processes$: Observable<Process[]> = this.processesSubject.asObservable();

  addProcess(process: Process): void {
    this.processes.push(process);
    this.processesSubject.next(this.processes);
  }

  handleConnectionMessage(conn_message: Process) {
      console.log(conn_message);
  }

  deleteProcess(processName: string): void {
    this.processes = this.processes.filter(process => process.name !== processName);
    this.processesSubject.next(this.processes);
  }

  reconnectProcess(processName: string): void {
    const process = this.processes.find(p => p.name === processName);
    if (process) {
      process.connected = true;
      this.processesSubject.next(this.processes);
    }
  }

  disconnectProcess(process: Process){
    process.connected = false;
  }

  toggleMonitoring(process: Process){
    process.monitoring = !process.monitoring;
    //api logic here send message to the backend to stop sending
  }
}
