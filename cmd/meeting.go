package cmd

import (
	"github.com/Huangscar/repo/goagenda-master/controller"
	"github.com/spf13/cobra"
)

var meetingRootCmd = &cobra.Command{
	Use:     "meeting",
	Aliases: []string{"m"},
	Short:   "Meeting management",
	Long:    "Meeting management",
}

var meetingDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete your meeting",
	Long:    "Delete your meeting",
	Run:     wrapper(controller.GetMeetingCtrl().DeleteMeeting),
}

var meetingAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add your meeting participator",
	Long:    "Add your meeting participator",
	Args:    cobra.MinimumNArgs(1),
	Run:     wrapper(controller.GetMeetingCtrl().Add),
}

var meetingListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List your meeting",
	Long:    "List your meeting",
	Run:     wrapper(controller.GetMeetingCtrl().ListMeeting),
}

var meetingRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"r"},
	Short:   "Remove your meeting participator",
	Long:    "Remove your meeting participator",
	Args:    cobra.MinimumNArgs(1),
	Run:     wrapper(controller.GetMeetingCtrl().Remove),
}

func init() {
	rootCmd.AddCommand(meetingRootCmd)
	meetingRootCmd.AddCommand(meetingDeleteCmd)
	meetingRootCmd.AddCommand(meetingAddCmd)
	meetingRootCmd.AddCommand(meetingListCmd)
	meetingRootCmd.AddCommand(meetingRemoveCmd)

	meetingListCmd.Flags().StringP("title", "t", "", "the meeting title")
	meetingListCmd.Flags().StringP("startTime", "s", "", "the meeting start time")
	meetingListCmd.Flags().StringP("endTime", "e", "", "the meeting end time")
	meetingListCmd.Flags().StringP("userName", "u", "", "the user name of the host of the meeting")
	meetingListCmd.MarkFlagRequired("title")

	meetingDeleteCmd.Flags().StringP("title", "t", "", "new password")
	meetingDeleteCmd.Flags().BoolP("all", "a", false, "delete all meeting")

	meetingAddCmd.Flags().StringP("title", "t", "", "meeting title")
	meetingAddCmd.MarkFlagRequired("title")

	meetingRemoveCmd.Flags().StringP("title", "t", "", "meeting title")
	meetingRemoveCmd.MarkFlagRequired("title")
}
