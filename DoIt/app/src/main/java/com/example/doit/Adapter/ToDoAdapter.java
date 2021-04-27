package com.example.doit.Adapter;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.CheckBox;

import com.example.doit.MainActivity;
import com.example.doit.Model.ToDoModel;
import com.example.doit.R;

import java.util.List;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

public class ToDoAdapter extends RecyclerView.Adapter<ToDoAdapter.ViewHolder> {

    private List<ToDoModel> todoList;  //lists of task
    private MainActivity activity;

    public ToDoAdapter(MainActivity activity)
    {
        this.activity=activity;
    }
    public static class ViewHolder extends RecyclerView.ViewHolder{
        CheckBox task;

        ViewHolder(View view)
        {
            super(view);
            task=view.findViewById(R.id.todoCheckBox);
        }

    }
    @NonNull
    @Override
    public ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View itemView = LayoutInflater.from(parent.getContext())
                .inflate(R.layout.task_layout, parent, false);
        return new ViewHolder(itemView);
    }


    private boolean toBoolean(int n)
    {
        if(n==1)return true;
        else return false;
    }

    public void setTasks(List<ToDoModel> todoList)
    {
        this.todoList=todoList;
        notifyDataSetChanged();
    }
    public void onBindViewHolder(ViewHolder holder,int position)
    {
        ToDoModel item =todoList.get(position);
        holder.task.setText(item.getTask());
        holder.task.setChecked(toBoolean(item.getStatus()));
    }

    public int getItemCount()
    {
        return todoList.size();
    }


}
