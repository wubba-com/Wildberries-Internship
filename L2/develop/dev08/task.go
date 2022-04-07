package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/**
8. Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд
*/

/**
Пакет exec запускает внешние команды. Он обертывает os.StartProcess,
чтобы сделать его проще переназначить stdin и stdout,
соединить ввод /вывод с помощью каналов и сделать другие корректировки.
*/

const (
	CommandEcho = "echo"
	CommandCd   = "cd"
	CommandKill = "kill"
	CommandPwd  = "pwd"
	CommandExit = "quit"
	CommandPs   = "ps"
	ExitText    = "Exit"
)

type Command interface {
	Exec(args ...string) ([]byte, error)
}

// echoCmd выполняет unix-команду echo и возвращает результат в байтах
type echoCmd struct {
}

func (e *echoCmd) Exec(args ...string) ([]byte, error) {
	return exec.Command("echo", args...).Output()
}

// cdCmd - изменяет директорию
type cdCmd struct {
}

func (c *cdCmd) Exec(args ...string) ([]byte, error) {
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		return nil, err
	}
	dir, err = os.Getwd()
	if err != nil {
		return nil, err
	}

	return []byte(dir), nil
}

// pwdCmd - выводит путь директории в которой находится терминал
type pwdCmd struct {
}

func (p *pwdCmd) Exec(args ...string) ([]byte, error) {
	dir, err := os.Getwd()
	return []byte(dir), err
}

// killCmd - убивает запущенный процесс
type killCmd struct {
}

func (k *killCmd) Exec(args ...string) ([]byte, error) {
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = process.Kill()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return []byte("killed"), nil
}

// psCmd - Выводит работающие процессы
type psCmd struct {
}

func (p *psCmd) Exec(args ...string) ([]byte, error) {
	return exec.Command("ps").Output()
}

// Shell - UNIX-шелл-утилита с поддержкой ряда простейших команд
type Shell struct {
	command Command
	output  io.Writer
}

func (s *Shell) SetCommand(cmd Command) {
	s.command = cmd
}

// Run - выполнение конкретной команды
func (s *Shell) run(args ...string) {
	b, err := s.command.Exec(args...)
	_, err = fmt.Fprintln(s.output, string(b))
	if err != nil {
		fmt.Println("[err]", err.Error())
		return
	}
}

// ExecuteCommands Исполняет команды, которые ввел пользователь
func (s *Shell) ExecuteCommands(cmds []string) {
	for _, command := range cmds {
		args := strings.Split(command, " ")

		com := args[0]
		if len(args) > 1 {
			args = args[1:]
		}

		switch com {
		case CommandEcho:
			cmd := &echoCmd{}
			s.SetCommand(cmd)

		case CommandCd:
			cmd := &cdCmd{}
			s.SetCommand(cmd)

		case CommandKill:
			cmd := &killCmd{}
			s.SetCommand(cmd)

		case CommandPwd:
			cmd := &pwdCmd{}
			s.SetCommand(cmd)

		case CommandPs:
			cmd := &psCmd{}
			s.SetCommand(cmd)

		case CommandExit: // завершение программы
			_, err := fmt.Fprintln(s.output, ExitText)
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}
			os.Exit(1)
		default:
			fmt.Println("Такой команды не будет")
			continue
		}
		s.run(args...)
	}
}

func main() {
	// Читает из стандартного ввода
	scan := bufio.NewScanner(os.Stdin)

	// устанавливается общий вывод результата команд
	var output = os.Stdout

	shell := &Shell{output: output}
	for {
		fmt.Print("command: ")

		if scan.Scan() {
			line := scan.Text()
			cmds := strings.Split(line, " | ")

			shell.ExecuteCommands(cmds)
		}
	}
}
