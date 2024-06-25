import { Injectable } from '@angular/core';
import { Log, SeverityLevel } from '../models/log';
import { BehaviorSubject, Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class LogService {

  private logs: Log[] = [
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
