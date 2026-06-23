// NewRenderer()
//   -> 解析所有页面模板
//   -> 存进 Renderer.pages
//   -> 返回 *Renderer

// renderer.HTML(...)
//   -> 根据 page 名字找模板
//   -> 找不到就报错
//   -> 找到就执行 layout 模板
//   -> 把 HTML 写回浏览器