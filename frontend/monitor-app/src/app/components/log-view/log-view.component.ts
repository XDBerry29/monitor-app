import { Component, ElementRef, HostListener, ViewChild, AfterViewChecked } from '@angular/core';
import {Log, SeverityLevel} from '../../models/log'
import { LogService } from '../../services/log.service';
import { Observable, Subscription, tap } from 'rxjs';

@Component({
  selector: 'app-log-view',
  templateUrl: './log-view.component.html',
  styleUrl: './log-view.component.css'
})
export class LogViewComponent implements AfterViewChecked{
  @ViewChild('logMessagesContainer', { static: false })
  logMessagesContainer!: ElementRef;

  logs$: Observable<Log[]>;
  ScrollToBottom = true;
  severityLevels = Object.keys(SeverityLevel).filter(key => isNaN(Number(key)));
  minSeverity = "DEBUG";
  constructor(private logService: LogService){
    this.logs$ = this.logService.logs$
    console.log(this.minSeverity)
  }

  ngAfterViewChecked(): void {
    if(this.ScrollToBottom){
      this.scrollToBottom();
    }
  }



  clearLogs(): void {
    this.logService.clearLogs();
  }

  addLog(): void {
    this.logService.addLog({ severity: SeverityLevel.WARNING, time: new Date(), process: 'Process4', message: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce nec pretium metus. Quisque congue leo in commodo laoreet. Ut placerat accumsan mi, vel sollicitudin ligula tincidunt pellentesque. Praesent non erat hendrerit, faucibus mi a, sagittis augue. Integer efficitur sapien sit amet mi facilisis pretium. Praesent aliquet elementum lorem nec dignissim. Donec vitae placerat magna, vitae blandit augue. Suspendisse non sem at lacus posuere pellentesque et eget velit. Aenean fringilla facilisis neque, nec elementum nulla finibus vel. Aliquam eros orci, semper ac feugiat a, suscipit in est." });

  }

  filterLogs() {
    console.log(this.minSeverity)
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
