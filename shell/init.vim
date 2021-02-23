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
"下面是对molokai的主题配置,需要将molokai.vim文件拷贝到/usr/share/vim/vim74/colors
set t_Co=256
colorscheme molokai
let g:rehash256 = 1 
let g:molokai_originl = 1	"1浅色，0深色 
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
" golang
Plug 'fatih/vim-go'
" commenter 注释
Plug 'scrooloose/nerdcommenter'
" coc
Plug 'neoclide/coc.nvim', {'branch': 'release'}
" rust
Plug 'rust-lang/rust.vim'
" luochen1990/rainbow 彩虹括号
Plug 'luochen1990/rainbow'

call plug#end()

"-------->>NERD Tree<<-------
"宏F2打开目录树
nmap <F2> :NERDTreeToggle <CR>
"-------->>NERD Commenter<<---
let mapleader=","
" ,cc 注释当前行
" ,cs 以性感方式注释
" ,cu 取消注释

"------>>> rust <<<---------
let g:rustfmt_autosave = 1

"------->>>rainbow 彩虹括号<<<-----
let g:rainbow_active = 1
