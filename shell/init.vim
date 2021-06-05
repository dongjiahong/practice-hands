"这里是我常用的一些vim的配置

syntax on			"开启语法高亮
set nu				"行号        
set hls				"高亮
set cursorline		"突出当前行
hi CursorLine cterm=NONE ctermbg=236 ctermfg=NONE
set ruler			"打开状态栏标尺

set tabstop=4		"tab键的宽度为4
set expandtab
set softtabstop=4	"使用退格键时，一次删除4个空格
set shiftwidth=4	"设置<<和>>命令移动时的宽度为4

set completeopt=menu "关闭scratch 

"set cindent			"c风格的换行
set backspace=2		"mac机器需要开这个能用backspace键
set nofoldenable	"不折叠
set autoindent		"自动换行

set colorcolumn=100  "100个字符为竖线
"set textwidth=100	"100个字符一行
set fo+=mB			"支持汉语

set smartindent		"开启新行时使用智能自动缩进
set nobackup		"不允许自动备份

set laststatus=2	"显示状态栏
set statusline=[%F]%y%r%m%*%=[Line:%l/%L,Column:%c][%p%%]
"XXX 对于主题
""Neovim:
""cd phanviet/colors
""mv monokai_pro.vim ~/.config/nvim/colors
""Vim:
""cd phanviet/colors
""mv monokai_pro.vim ~/.vim/colors
" ==========>  molokai <========
"set t_Co=256
"colorscheme molokai
let g:rehash256 = 1 
"let g:molokai_originl = 1	"1浅色，0深色 
" ==========> monokai_pro <=======
set termguicolors
"colorscheme chito
"colorscheme miramare
colorscheme monokai_pro
" 打开时光标在上次的位置
au BufReadPost * if line("'\"") > 0|if line("'\"") <= line("$")|exe("norm'\"")|else|exe "norm $"|endif|endif 


" vim-plug插件管理器的安装
" 执行命令： curl -fLo ~/.config/nvim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

filetype plugin indent on
syntax on
call plug#begin('~/.vim/plugged')

" nerd tree
Plug 'scrooloose/nerdtree'
" molokai colors
Plug 'tomasr/molokai'
Plug 'Jimeno0/vim-chito'
Plug 'phanviet/vim-monokai-pro'
Plug 'franbach/miramare'
" golang
Plug 'fatih/vim-go'
" gotags
Plug 'jstemmer/gotags'
Plug 'preservim/tagbar'
" commenter 注释
Plug 'scrooloose/nerdcommenter'
" coc
Plug 'neoclide/coc.nvim', {'branch': 'release'}
" rust
" luochen1990/rainbow 彩虹括号
Plug 'luochen1990/rainbow'
" toml
Plug 'cespare/vim-toml'

call plug#end()

"-------->>NERD Tree<<-------
"宏F2打开目录树
nmap <F2> :NERDTreeToggle <CR>
"-------->>NERD Commenter<<---
let mapleader=","
" ,cc 注释当前行
" ,cs 以性感方式注释
" ,cu 取消注释

"------>>> coc <<<---------
" 设置悬浮文档
nnoremap <silent> <leader>h :call CocActionAsync('doHover')<CR>
" 批量重命名
nmap <leader>r <Plug>(coc-rename)
" 重构：把光标下的变量/trait/类型所有相关的代码提取出来，在左窗口统一修改
nmap <leader>rf <Plug>(coc-refactor)
nmap <leader>cr <Plug>(coc-references)

"------->>>rainbow 彩虹括号<<<-----
let g:rainbow_active = 1

" ---------> vim-go <-------------
let g:go_fmt_autosave=0
let g:go_fmt_command="gofmt"
let g:go_imports_autosave=0 " 保存时不自动导入包--太慢了
"let g:go_template_autocreate = 0 " 新文件不自动填充
let g:go_highlight_types = 1
let g:go_highlight_fields = 1 " 变量高亮
let g:go_highlight_functions = 1 " 函数高亮
let g:go_highlight_function_calls = 1
let g:go_highlight_operators = 1
let g:go_highlight_extra_types = 1
let g:go_highlight_build_constraints = 1
let g:go_highlight_generate_tags = 1
let g:go_code_completion_enabled = 0
" dsable all linters as that is taken care of by coc.nvim
let g:go_diagnostics_enabled = 0
let g:go_metalinter_enabled = []
"Find all references of a given type/function in the codebase with ,+cr:
"autocmd BufEnter *.go nmap <leader>cr <Plug>(coc-references)
"Not many options here, but there’s renaming the symbol your cursor is on with ,+r:
"autocmd BufEnter *.go nmap <leader>r <Plug>(coc-rename)
" 不让coc-go提示
"autocmd FileType go let b:coc_suggest_disable = 1
"autocmd VimEnter *.go NERDTreeToggle

" -------------> gotags <------------
map <F8> :Tagbar<CR>
let g:tagbar_type_go = {
    \ 'ctagstype' : 'go',
    \ 'kinds'     : [
        \ 'p:package',
        \ 'i:imports:1',
        \ 'c:constants',
        \ 'v:variables',
        \ 't:types',
        \ 'n:interfaces',
        \ 'w:fields',
        \ 'e:embedded',
        \ 'm:methods',
        \ 'r:constructor',
        \ 'f:functions'
    \ ],
    \ 'sro' : '.',
    \ 'kind2scope' : {
        \ 't' : 'ctype',
        \ 'n' : 'ntype'
    \ },
    \ 'scope2kind' : {
        \ 'ctype' : 't',
        \ 'ntype' : 'n'
    \ },
    \ 'ctagsbin'  : 'gotags',
    \ 'ctagsargs' : '-sort -silent'
\ }
