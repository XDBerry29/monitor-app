import { Component, Input } from '@angular/core';
import { Log, SeverityLevel } from '../../models/log';

@Component({
  selector: 'app-log-message',
  templateUrl: './log-message.component.html',
  styleUrl: './log-message.component.css'
})
export class LogMessageComponent {
  @Input()
  log!: Log;
  expanded: boolean = false;

  getSeverityName(severity: SeverityLevel): string {
    return SeverityLevel[severity];
  }

  toggleDetails(): void {
    this.expanded = !this.expanded;
  }

  getClasses() {
    return {
      'expanded': this.expanded,
      'log-message': true,
      [this.getSeverityName(this.log.severity)]: true
    };
  }

  
}
