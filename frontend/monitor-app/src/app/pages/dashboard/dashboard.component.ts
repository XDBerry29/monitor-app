import { Component } from '@angular/core';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.css'
})
export class DashboardComponent {

  currentMode='dark-mode'
  
  toggleMode() {
    if(this.currentMode == 'dark-mode'){
      this.currentMode='light-mode'
    }else {
      this.currentMode ='dark-mode'
    }
  }
    

}
