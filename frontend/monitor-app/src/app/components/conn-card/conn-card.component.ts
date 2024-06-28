import { ProcessFilterMsg } from './../../models/process-filter-msg';
import { SocketService } from './../../services/socket.service';
import { Component, Input } from '@angular/core';
import { Process } from '../../models/process';
import { ProcessesService } from '../../services/processes.service';

@Component({
  selector: 'app-conn-card',
  templateUrl: './conn-card.component.html',
  styleUrl: './conn-card.component.css'
})
export class ConnCardComponent {

  @Input() process!: Process;

  constructor(private service: ProcessesService, private socketService:SocketService) { }

  onToggleMonitoring(): void {
    this.service.toggleMonitoring(this.process);
    let filterMsg : ProcessFilterMsg = {
      name: this.process.name,
      monitoring: this.process.monitoring
    }
    this.socketService.sendProcessFilterMessage(filterMsg);
  }

  removeProccesFromList() {
    this.service.deleteProcess(this.process.name);
  }
}
