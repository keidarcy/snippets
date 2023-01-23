package com.xyh.basic;

public abstract class UIControl {
    boolean isEnable = true;

//    public UIControl(boolean isEnable) {
//        this.isEnable = isEnable;
////        System.out.println("UIControl");
//    }

    public abstract void render();

    public void enable() {
        isEnable = true;
    }

    public void disable() {
        isEnable = false;
    }

    public boolean isEnable() {
        return isEnable;
    }
}
