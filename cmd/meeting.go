package cmd

import (
	"github.com/MegaShow/goagenda/controller"
	"github.com/spf13/cobra"
)

var meetingRootCmd = &cobra.Command{
	Use:     "meeting",
	Aliases: []string{"m", "meet"},
	Short:   "Meeting management",
	Long:    "Meeting management",
}

var meetingCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create a meeting",
	Long:    "Create a meeting",
	Run:     wrapper(controller.GetMeetingCtrl().MeetingCreate),
}

var meetingSetCmd = &cobra.Command{
	Use:     "set",
	Aliases: []string{"s"},
	Short:   "Set information of a meeting",
	Long:    "Set information of a meeting",
	Run:     wrapper(controller.GetMeetingCtrl().MeetingSet),
}

var meetingQuitCmd = &cobra.Command{
	Use:     "quit",
	Aliases: []string{"q"},
	Short:   "Quit a meeting",
	Long:    "Quit a meeting",
	Run:     wrapper(controller.GetMeetingCtrl().MeetingQuit),
}

func init() {
	rootCmd.AddCommand(meetingRootCmd)
	meetingRootCmd.AddCommand(meetingCreateCmd)
	meetingRootCmd.AddCommand(meetingSetCmd)
	meetingRootCmd.AddCommand(meetingQuitCmd)

	meetingCreateCmd.Flags().StringP("title", "t", "", "title of new meeting")
	meetingCreateCmd.Flags().StringP("startTime", "s", "", "start time of new meeting")
	meetingCreateCmd.Flags().StringP("endTime", "e", "", "end time of new meeting")
	meetingCreateCmd.Flags().StringSliceP("participator", "p", []string{}, "participators of new meeting")
	meetingCreateCmd.MarkFlagRequired("title")
	meetingCreateCmd.MarkFlagRequired("startTime")
	meetingCreateCmd.MarkFlagRequired("endTime")
	meetingCreateCmd.MarkFlagRequired("participator")

	meetingSetCmd.Flags().StringP("title", "t", "", "title of new meeting")
	meetingSetCmd.Flags().StringP("startTime", "s", "", "start time of new meeting")
	meetingSetCmd.Flags().StringP("endTime", "e", "", "end time of new meeting")
	meetingSetCmd.Flags().StringSliceP("participator", "p", []string{}, "participators of new meeting")
	meetingSetCmd.MarkFlagRequired("title")

	meetingQuitCmd.Flags().StringP("title", "t", "", "title of new meeting")
	meetingQuitCmd.MarkFlagRequired("title")
}
