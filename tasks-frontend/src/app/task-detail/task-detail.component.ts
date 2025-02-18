import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';
import { TaskService } from '../services/task.service';
import { Task } from '../models/task';
import { CommonModule } from '@angular/common';

@Component({
  imports: [CommonModule, RouterModule],
  selector: 'app-task-detail',
  templateUrl: './task-detail.component.html',
  styleUrls: ['./task-detail.component.scss']
})
export class TaskDetailComponent implements OnInit {
  task: Task | null = null; 
  errorMessage: string = '';
  isLoading = true;

  constructor(
    private route: ActivatedRoute,
    private taskService: TaskService,
    private router: Router
  ) {}

  ngOnInit() {
    const id = Number(this.route.snapshot.paramMap.get('id'));

    this.taskService.getTaskById(id).subscribe({
      next: (data) => {
        if (data) {
          this.task = data;
        } else {
          this.errorMessage = 'Task not found.';
        }
        this.isLoading = false;
      },
      error: (err) => {
        console.error('There was an error fetching the task:', err);
        this.errorMessage = 'An error occurred while fetching the task.';
        this.isLoading = false;
      }
    });
  }

  goBack() {
    this.router.navigate(['/']);
  }
}
