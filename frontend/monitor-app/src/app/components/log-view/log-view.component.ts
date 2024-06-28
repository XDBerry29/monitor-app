import { LogFilterMsg } from './../../models/log-filter-msg';
import { Component, ElementRef, HostListener, ViewChild, AfterViewChecked, OnInit } from '@angular/core';
import {Log, SeverityLevel} from '../../models/log'
import { LogService } from '../../services/log.service';
import { Observable, Subscription, tap } from 'rxjs';
import { SocketService } from '../../services/socket.service';

@Component({
  selector: 'app-log-view',
  templateUrl: './log-view.component.html',
  styleUrl: './log-view.component.css'
})
export class LogViewComponent implements AfterViewChecked{
  @ViewChild('logMessagesContainer', { static: false })
  logMessagesContainer!: ElementRef;

  logs$: Observable<Log[]>;
  dropdownOpen = false;
  ScrollToBottom = true;
  severityLevels = Object.keys(SeverityLevel).filter(key => isNaN(Number(key)));
  displayedSeverityLevels: string[] = [...this.severityLevels];
  minSeverity = 'DEBUG';
  constructor(private logService: LogService, private wsService: SocketService){
    this.logs$ = this.logService.logs$
    this.filterLogs()
  }

  ngAfterViewChecked(): void {
    if(this.ScrollToBottom){
      this.scrollToBottom();
    }
  }

  toggleDropdown(): void {
    if (this.dropdownOpen) {
      this.dropdownOpen = false;
    } else {
      this.dropdownOpen = true;
    }
  }

  selectOption(value: string): void {
    this.minSeverity = value;
    this.toggleDropdown()
    this.filterLogs();
  }



  clearLogs(): void {
    this.logService.clearLogs();
    this.ScrollToBottom = true;
  }

  addLog(): void {
    this.logService.addLog({ severity: SeverityLevel.WARNING, time: "12:00:00", process: 'Process4', message: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce nec pretium metus. Quisque congue leo in commodo laoreet. Ut placerat accumsan mi, vel sollicitudin ligula tincidunt pellentesque. Praesent non erat hendrerit, faucibus mi a, sagittis augue. Integer efficitur sapien sit amet mi facilisis pretium. Praesent aliquet elementum lorem nec dignissim. Donec vitae placerat magna, vitae blandit augue. Suspendisse non sem at lacus posuere pellentesque et eget velit. Aenean fringilla facilisis neque, nec elementum nulla finibus vel. Aliquam eros orci, semper ac feugiat a, suscipit in est." });

  }

  filterLogs() {
    console.log(SeverityLevel[this.minSeverity as keyof typeof SeverityLevel])
    console.log(this.minSeverity)
    let filterMsg : LogFilterMsg = {
      severity: SeverityLevel[this.minSeverity as keyof typeof SeverityLevel]
    }
    this.wsService.sendLogFilterMessage(filterMsg)
  }

  scrollToBottom() {
    if (this.logMessagesContainer?.nativeElement) {
      this.logMessagesContainer.nativeElement.scrollTop = this.logMessagesContainer.nativeElement.scrollHeight ;

    }
  }

  @HostListener('scroll', ['$event'])
  onScroll(event: any) {
      if (event.target.offsetHeight + event.target.scrollTop >= event.target.scrollHeight) {
        this.ScrollToBottom = true;
      } else {
        this.ScrollToBottom = false;
      }
  }

}
