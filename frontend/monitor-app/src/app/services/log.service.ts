import { Injectable } from '@angular/core';
import { Log, SeverityLevel } from '../models/log';
import { BehaviorSubject, Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class LogService {

  private logs: Log[] = [
    { severity: SeverityLevel.DEBUG, time: new Date(), process: 'Process3', message: 'Debug message' },
    { severity: SeverityLevel.INFO, time: new Date(), process: 'Process1', message: 'Info message' },
    { severity: SeverityLevel.WARNING, time: new Date(), process: 'Process4', message: 'Warning message' },
    { severity: SeverityLevel.ERROR, time: new Date(), process: 'Process2', message: 'Error message' },
    { severity: SeverityLevel.CRITICAL, time: new Date(), process: 'Process2', message: 'Critical message' }
  ];

  private logsSource = new BehaviorSubject<Log[]>(this.logs);
  logs$ = this.logsSource.asObservable();

  constructor() { }

  // Add a log to the list
  addLog(log: Log): void {
    this.logs.push(log);
    this.logsSource.next(this.logs); // Update BehaviorSubject with the new reference
  }

  // Clear the list of logs
  clearLogs(): void {
    this.logs = [];
    this.logsSource.next(this.logs); // Update BehaviorSubject with an empty array
  }
      
  
   
}
