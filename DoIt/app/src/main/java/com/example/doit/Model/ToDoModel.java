package com.example.doit.Model;

public class ToDoModel {
    private int id,status;  //corresponding to one task of the list
    private String task;
    //getter setter methods...inbuilt
    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getStatus() {
        return status;
    }

    public void setStatus(int status) {
        this.status = status;
    }

    public String getTask() {
        return task;
    }

    public void setTask(String task) {
        this.task = task;
    }
}
