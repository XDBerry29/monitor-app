import { Injectable } from '@angular/core';
import { Process } from '../models/process';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ProcessesService {
  private processes: Process[] = [
    { name: 'API', connected: false, monitoring: false },
    { name: 'Web', connected: true, monitoring: true },
    { name: 'Worker', connected: true, monitoring: true },
    { name: 'Scheduler', connected: true, monitoring: true },
    { name: 'Daemon', connected: true, monitoring: true },
    { name: 'Cron', connected: true, monitoring: true },
    { name: 'All', connected: true, monitoring: true },{ name: 'All', connected: true, monitoring: true },{ name: 'All', connected: true, monitoring: true },{ name: 'All', connected: true, monitoring: true },{ name: 'All', connected: true, monitoring: true },{ name: 'All', connected: true, monitoring: true },{ name: 'All', connected: true, monitoring: true },{ name: 'All', connected: true, monitoring: true },{ name: 'All', connected: true, monitoring: true },{ name: 'All', connected: true, monitoring: true },{ name: 'All', connected: true, monitoring: true },{ name: 'All', connected: true, monitoring: true },
  ];

  private processesSubject = new BehaviorSubject<Process[]>(this.processes);
  processes$: Observable<Process[]> = this.processesSubject.asObservable();

  addProcess(process: Process): void {
    this.processes.push(process);
    this.processesSubject.next(this.processes);
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
