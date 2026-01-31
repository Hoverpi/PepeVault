const std = @import("std");
const sdl = @cImport({
    @cInclude("SDL3/SDL.h");
    @cInclude("SDL3/SDL_render.h");
    @cInclude("SDL3/SDL_mouse.h");
    @cInclude("SDL3/SDL_error.h");
    @cInclude("SDL3/SDL_events.h");
    @cInclude("SDL3/SDL_main.h");
    @cInclude("SDL3_ttf/SDL_ttf.h");
});

const WindowOptions = struct {
    title: ?[*:0]const u8 = null,
    width: ?i32 = null,
    height: ?i32 = null,
    flags: ?sdl.SDL_WindowFlags = null,
    renderer_name: ?[*:0]const u8 = null,
};

const Window = struct {
    window: ?*sdl.SDL_Window = null,
    w_width: i32 = 0,
    w_height: i32 = 0,
    w_flags: sdl.SDL_WindowFlags = 0,

    renderer: ?*sdl.SDL_Renderer = null,

    pub fn init(self: *Window, opts: *const WindowOptions) !void {
        self.window = sdl.SDL_CreateWindow(opts.title orelse null, opts.width orelse 0, opts.height orelse 0, opts.flags orelse sdl.SDL_WINDOW_RESIZABLE) orelse {
            std.log.err("SDL Error: {s}\n", .{sdl.SDL_GetError()});
            return;
        };
        self.renderer = sdl.SDL_CreateRenderer(self.window, opts.renderer_name orelse null) orelse {
            std.log.err("SDL Error: {s}\n", .{sdl.SDL_GetError()});
            return;
        };
        return;
    }

    pub fn destroy(self: *Window) void {
        if (self.renderer) |ren| {
            sdl.SDL_DestroyRenderer(ren);
            self.renderer = null;
        }
        if (self.window) |win| {
            sdl.SDL_DestroyWindow(win);
            self.window = null;
        }
    }

    pub fn maximizeWindow(self: *Window) void {
        if (self.window) |win| {
            _ = sdl.SDL_ShowWindow(win);
            _ = sdl.SDL_MaximizeWindow(win);
        }
    }

    pub fn getSurface(self: *Window) ?*sdl.SDL_Surface {
        if (self.window) |win| {
            sdl.SDL_GetWindowSurface(win);
        } else {
            return null;
        }
    }

    pub fn render(self: *Window) void {
        if (self.renderer) |ren| {
            _ = sdl.SDL_SetRenderDrawColor(ren, 170, 170, 170, 255);
            _ = sdl.SDL_RenderClear(ren);

            // update frame
            _ = sdl.SDL_RenderPresent(ren);
        }
    }

    pub fn getMouse(self: *Window) struct { mouse_x: f32, mouse_y: f32 } {
        var x: f32 = 0;
        var y: f32 = 0;

        if (self.window) |_| {
            _ = sdl.SDL_GetMouseState(&x, &y);
        }

        return .{ .mouse_x = x, .mouse_y = y };
    }
};

pub fn main() !void {
    if (!sdl.SDL_Init(sdl.SDL_INIT_VIDEO)) {
        std.log.err("SDL Error: {s}\n", .{sdl.SDL_GetError()});
        return;
    }
    defer sdl.SDL_Quit();

    var window: Window = .{};
    try window.init(&WindowOptions{
        .title = "PepeVault",
        .width = 600,
        .height = 400,
        .flags = sdl.SDL_WINDOW_RESIZABLE,
        .renderer_name = null,
    });
    defer window.destroy();

    window.maximizeWindow();

    var event: sdl.SDL_Event = undefined;
    var is_running: bool = true;

    var mouse = window.getMouse();

    while (is_running) {
        while (sdl.SDL_PollEvent(&event)) {
            switch (event.type) {
                sdl.SDL_EVENT_QUIT => is_running = false,
                else => {},
            }
        }
        mouse = window.getMouse();
        std.debug.print("({}, {})\n", .{ mouse.mouse_x, mouse.mouse_y });

        window.render();
    }
}
