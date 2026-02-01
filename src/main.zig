const std = @import("std");
const sdl = @cImport({
    @cInclude("SDL3/SDL.h");
    @cInclude("SDL3/SDL_render.h");
    @cInclude("SDL3/SDL_mouse.h");
    @cInclude("SDL3/SDL_pixels.h");
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

    pub fn init(self: *Window, opts: WindowOptions) !void {
        const title: [*:0]const u8 = opts.title orelse "App";
        const width: i32 = opts.width orelse 800;
        const height: i32 = opts.height orelse 600;
        const flags: sdl.SDL_WindowFlags = opts.flags orelse sdl.SDL_WINDOW_RESIZABLE;

        const win: ?*sdl.SDL_Window = sdl.SDL_CreateWindow(title, width, height, flags);
        if (win == null) {
            std.log.err("SDL Error: {s}\n", .{sdl.SDL_GetError()});
            return;
        }
        self.window = win;

        const ren: ?*sdl.SDL_Renderer = sdl.SDL_CreateRenderer(self.window, opts.renderer_name orelse null);
        if (ren == null) {
            std.log.err("SDL Error: {s}\n", .{sdl.SDL_GetError()});
            return;
        }
        self.renderer = ren;

        self.w_width = width;
        self.w_height = height;
        self.w_flags = flags;

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
            return sdl.SDL_GetWindowSurface(win);
        } else {
            return null;
        }
    }

    pub fn getRenderer(self: *Window) ?*sdl.SDL_Renderer {
        if (self.window) |win| {
            return sdl.SDL_GetRenderer(win);
        } else {
            return null;
        }
    }

    pub fn update(self: *Window) void {
        if (self.renderer) |ren| {
            _ = sdl.SDL_SetRenderDrawColor(ren, 170, 170, 170, 255);
            _ = sdl.SDL_RenderClear(ren);
        }
    }

    pub fn render(self: *Window) void {
        if (self.renderer) |ren| {
            _ = sdl.SDL_RenderPresent(ren);
        }
    }

    pub fn getMouse(self: *Window) struct { mouse_x: f32, mouse_y: f32 } {
        var x: f32 = 0;
        var y: f32 = 0;
        if (self.window) |_| _ = sdl.SDL_GetMouseState(&x, &y);
        return .{ .mouse_x = x, .mouse_y = y };
    }
};

const RectangleOptions = struct {
    rectangle: ?sdl.SDL_FRect = null,
    color: ?sdl.SDL_Color = null,
    enable_hover: ?bool = null,
    is_outline: ?bool = null,
    thickness: ?f32 = null,
};

const Rectangle = struct {
    rect: sdl.SDL_FRect = .{},
    base_color: sdl.SDL_Color = .{},
    current_color: sdl.SDL_Color = .{},
    can_hover: bool = false,
    is_hovered: bool = false,
    is_outline: bool = false,
    thickness: f32 = 1.0,

    pub fn init(self: *Rectangle, opts: RectangleOptions) void {
        self.rect = opts.rectangle orelse .{ .x = 0, .y = 0, .w = 100, .h = 100 };
        self.base_color = opts.color orelse .{ .r = 255, .g = 255, .b = 255, .a = 255 };
        self.current_color = self.base_color;
        self.can_hover = opts.enable_hover orelse false;
        self.is_outline = opts.is_outline orelse false;
        self.thickness = opts.thickness orelse 1.0;
    }

    pub fn updateHover(self: *Rectangle, mouse_x: f32, mouse_y: f32) void {
        if (!self.can_hover) return;

        const mouse_point: sdl.SDL_FPoint = .{ .x = mouse_x, .y = mouse_y };
        const inside = sdl.SDL_PointInRectFloat(&mouse_point, &self.rect);

        if (inside) {
            self.is_hovered = true;
            self.current_color = .{ .r = 0, .g = 255, .b = 0, .a = 255 };
        } else {
            self.is_hovered = false;
            self.current_color = self.base_color;
        }
    }

    pub fn render(self: *Rectangle, renderer: *sdl.SDL_Renderer) void {
        _ = sdl.SDL_SetRenderDrawColor(renderer, self.current_color.r, self.current_color.g, self.current_color.b, self.current_color.a);

        if (self.is_outline) {
            var i: f32 = 0;
            while (i < self.thickness) : (i += 1) {
                const thick_rect = sdl.SDL_FRect{
                    .x = self.rect.x + i,
                    .y = self.rect.y + i,
                    .w = self.rect.w - (i * 2),
                    .h = self.rect.h - (i * 2),
                };
                _ = sdl.SDL_RenderRect(renderer, &thick_rect);
            }
        } else {
            _ = sdl.SDL_RenderFillRect(renderer, &self.rect);
        }
    }

    pub fn setPosition(self: *Rectangle, x: f32, y: f32) void {
        self.rect.x = x;
        self.rect.y = y;
    }

    pub fn getRect(self: *Rectangle) sdl.SDL_FRect {
        return self.rect;
    }
};

pub fn main() !void {
    if (!sdl.SDL_Init(sdl.SDL_INIT_VIDEO)) {
        std.log.err("SDL Error: {s}\n", .{sdl.SDL_GetError()});
        return;
    }
    defer sdl.SDL_Quit();

    var window: Window = .{};
    try window.init(.{
        .title = "PepeVault",
        .width = 600,
        .height = 400,
        .flags = sdl.SDL_WINDOW_RESIZABLE,
        .renderer_name = null,
    });
    defer window.destroy();

    window.maximizeWindow();

    var margin: Rectangle = .{};
    margin.init(.{
        .rectangle = .{ .x = 0, .y = 0, .w = 400, .h = 400 },
        .color = .{ .r = 91, .g = 96, .b = 181, .a = 255 },
        .enable_hover = false,
        .is_outline = true,
        .thickness = 10.0,
    });

    var login_button: Rectangle = .{};
    login_button.init(.{
        .rectangle = .{ .x = 0, .y = 0, .w = 100, .h = 40 },
        .color = .{ .r = 255, .g = 0, .b = 0, .a = 255 },
        .enable_hover = true,
        .is_outline = false,
    });

    var event: sdl.SDL_Event = undefined;
    var is_running: bool = true;

    while (is_running) {
        while (sdl.SDL_PollEvent(&event)) {
            switch (event.type) {
                sdl.SDL_EVENT_QUIT => is_running = false,

                sdl.SDL_EVENT_WINDOW_RESIZED => {
                    const win_w: f32 = @as(f32, @floatFromInt(event.window.data1));
                    const win_h: f32 = @as(f32, @floatFromInt(event.window.data2));

                    // center margin
                    const margin_size: sdl.SDL_FRect = margin.getRect();
                    margin.setPosition((win_w - margin_size.w) / 2.0, (win_h - margin_size.h) / 2.0);

                    // center button
                    const login_button_size: sdl.SDL_FRect = login_button.getRect();
                    login_button.setPosition((win_w - login_button_size.w) / 2.0, (win_h - login_button_size.h) / 2.0);
                },
                else => {},
            }
        }
        const mouse = window.getMouse();
        std.debug.print("({}, {})\n", .{ mouse.mouse_x, mouse.mouse_y });

        login_button.updateHover(mouse.mouse_x, mouse.mouse_y);

        window.update();
        if (window.getRenderer()) |ren| {
            margin.render(ren);

            const margin_rect: sdl.SDL_FRect = margin.getRect();
            const t = margin.thickness;
            const clip_rect: sdl.SDL_Rect = .{
                .x = @intFromFloat(margin_rect.x + t),
                .y = @intFromFloat(margin_rect.y + t),
                .w = @intFromFloat(margin_rect.w - (t * 2)),
                .h = @intFromFloat(margin_rect.h - (t * 2)),
            };
            _ = sdl.SDL_SetRenderClipRect(ren, &clip_rect);

            login_button.render(ren);

            _ = sdl.SDL_SetRenderClipRect(ren, null);
        }
        window.render();
    }
}
