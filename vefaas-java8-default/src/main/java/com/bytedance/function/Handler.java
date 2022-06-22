package com.bytedance.function;

import com.bytedance.bytefaas.context.Context;
import com.bytedance.bytefaas.handler.HttpServletHandler;
import com.bytedance.bytefaas.handler.Initializer;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

public class Handler implements HttpServletHandler, Initializer {
    private String msg = "Hello FaaS";

    public void handle(Context context, HttpServletRequest httpServletRequest, HttpServletResponse httpServletResponse) throws Exception {
        try {
            httpServletResponse.getWriter().println(msg);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    @Override
    public void initialize(Context context) throws Exception {
        msg = "Hello, ByteFaaS";
    }
}
